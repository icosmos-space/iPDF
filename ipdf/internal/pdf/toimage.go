package pdf

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"

	"ipdf/internal/job"
)

// ToImages rasterizes each PDF page to an image via go-fitz (MuPDF).
// format: jpg | png; dpi defaults to 144.
// Build with CGO so the bundled MuPDF static libs are linked into the binary.
func ToImages(path, format string, dpi int, outDir string) (ToolResult, error) {
	if err := validatePDFInput(path); err != nil {
		return ToolResult{}, err
	}
	format = strings.ToLower(strings.TrimSpace(format))
	if format == "jpeg" {
		format = "jpg"
	}
	if format != "jpg" && format != "png" {
		return ToolResult{}, fmt.Errorf("不支持的格式: %s（支持 jpg / png）", format)
	}
	if dpi <= 0 {
		dpi = 144
	}
	if outDir == "" {
		return ToolResult{}, fmt.Errorf("请指定输出目录")
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return ToolResult{}, err
	}

	ws, err := job.New()
	if err != nil {
		return ToolResult{}, err
	}
	defer ws.Cleanup()

	doc, err := fitz.New(path)
	if err != nil {
		return ToolResult{}, fmt.Errorf("打开 PDF 失败: %w", err)
	}
	defer doc.Close()

	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	tmpDir := ws.Path("pages")
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		return ToolResult{}, err
	}

	n := doc.NumPage()
	var outs []string
	for i := 0; i < n; i++ {
		img, err := doc.ImageDPI(i, float64(dpi))
		if err != nil {
			return ToolResult{}, fmt.Errorf("渲染第 %d 页失败: %w", i+1, err)
		}

		name := fmt.Sprintf("%s_p%03d.%s", base, i+1, format)
		tmpPath := filepath.Join(tmpDir, name)
		f, err := os.Create(tmpPath)
		if err != nil {
			return ToolResult{}, err
		}
		switch format {
		case "png":
			err = png.Encode(f, img)
		default:
			err = jpeg.Encode(f, img, &jpeg.Options{Quality: 90})
		}
		_ = f.Close()
		if err != nil {
			return ToolResult{}, fmt.Errorf("编码第 %d 页失败: %w", i+1, err)
		}

		dst := filepath.Join(outDir, name)
		if err := job.MoveOrCopy(tmpPath, dst); err != nil {
			return ToolResult{}, err
		}
		outs = append(outs, dst)
	}

	return ToolResult{
		OutputPaths: outs,
		Message:     fmt.Sprintf("已导出 %d 张图片到 %s", len(outs), outDir),
	}, nil
}
