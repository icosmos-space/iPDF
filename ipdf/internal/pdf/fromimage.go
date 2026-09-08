package pdf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"ipdf/internal/job"
)

var imageExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".tif": true, ".tiff": true,
}

// FromImages builds a PDF from image files (one page per image).
func FromImages(paths []string, outPath string) (ToolResult, error) {
	if len(paths) == 0 {
		return ToolResult{}, fmt.Errorf("请至少选择一张图片")
	}
	for _, p := range paths {
		ext := strings.ToLower(filepath.Ext(p))
		if !imageExts[ext] {
			return ToolResult{}, fmt.Errorf("不支持的图片格式: %s", p)
		}
		if err := fileReadable(p); err != nil {
			return ToolResult{}, err
		}
	}
	if outPath == "" {
		return ToolResult{}, fmt.Errorf("请指定输出路径")
	}

	ws, err := job.New()
	if err != nil {
		return ToolResult{}, err
	}
	defer ws.Cleanup()

	tmpOut := ws.Path("from_images.pdf")
	_ = os.Remove(tmpOut)

	imp := api.DefaultImportConfig()
	if err := api.ImportImagesFile(paths, tmpOut, imp, conf()); err != nil {
		return ToolResult{}, fmt.Errorf("图片转 PDF 失败: %w", err)
	}
	if err := job.MoveOrCopy(tmpOut, outPath); err != nil {
		return ToolResult{}, fmt.Errorf("写入输出文件失败: %w", err)
	}

	return ToolResult{
		OutputPaths: []string{outPath},
		Message:     fmt.Sprintf("已将 %d 张图片写入 %s", len(paths), filepath.Base(outPath)),
	}, nil
}
