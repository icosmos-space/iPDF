package pdf

import (
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	"golang.org/x/image/webp"

	"ipdf/internal/job"
)

// ToImages extracts embedded images from a PDF via pdfcpu.
// Note: pdfcpu cannot rasterize full pages; this exports image XObjects.
// format: jpg | png (converts when needed); dpi is accepted for API compat but unused.
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
	_ = dpi
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

	tmpDir := ws.Path("images")
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		return ToolResult{}, err
	}

	if err := api.ExtractImagesFile(path, tmpDir, nil, conf()); err != nil {
		return ToolResult{}, fmt.Errorf("提取图片失败: %w", err)
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		return ToolResult{}, err
	}

	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	var outs []string
	idx := 0
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		src := filepath.Join(tmpDir, e.Name())
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if !isImageExt(ext) {
			continue
		}
		idx++
		outName := fmt.Sprintf("%s_img%03d.%s", base, idx, format)
		dst := filepath.Join(outDir, outName)
		if err := convertOrCopyImage(src, dst, format); err != nil {
			return ToolResult{}, fmt.Errorf("处理 %s 失败: %w", e.Name(), err)
		}
		outs = append(outs, dst)
	}

	if len(outs) == 0 {
		return ToolResult{}, fmt.Errorf("未找到可提取的嵌入图片（pdfcpu 无法将文字页栅格化为图片；扫描件/图片型 PDF 通常可提取）")
	}

	return ToolResult{
		OutputPaths: outs,
		Message:     fmt.Sprintf("已提取 %d 张嵌入图片到 %s", len(outs), outDir),
	}, nil
}

func isImageExt(ext string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp", ".tif", ".tiff", ".gif", ".bmp":
		return true
	default:
		return false
	}
}

func convertOrCopyImage(src, dst, format string) error {
	srcExt := strings.ToLower(filepath.Ext(src))
	if (srcExt == ".jpg" || srcExt == ".jpeg") && format == "jpg" {
		return job.CopyFile(src, dst)
	}
	if srcExt == ".png" && format == "png" {
		return job.CopyFile(src, dst)
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	var img image.Image
	switch srcExt {
	case ".png":
		img, err = png.Decode(f)
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(f)
	case ".webp":
		img, err = webp.Decode(f)
	default:
		img, _, err = image.Decode(f)
	}
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	switch format {
	case "png":
		return png.Encode(out, img)
	default:
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 92})
	}
}
