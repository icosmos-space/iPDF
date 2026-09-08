package pdf

import (
	"fmt"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"ipdf/internal/job"
)

// Merge concatenates PDF files in order into outPath.
func Merge(paths []string, outPath string) (ToolResult, error) {
	if len(paths) < 2 {
		return ToolResult{}, fmt.Errorf("至少需要 2 个 PDF 文件")
	}
	for _, p := range paths {
		if err := validatePDFInput(p); err != nil {
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

	tmpOut := ws.Path("merged.pdf")
	if err := api.MergeCreateFile(paths, tmpOut, false, conf()); err != nil {
		return ToolResult{}, fmt.Errorf("合并失败: %w", err)
	}
	if err := job.MoveOrCopy(tmpOut, outPath); err != nil {
		return ToolResult{}, fmt.Errorf("写入输出文件失败: %w", err)
	}

	return ToolResult{
		OutputPaths: []string{outPath},
		Message:     fmt.Sprintf("已合并 %d 个文件 → %s", len(paths), filepath.Base(outPath)),
	}, nil
}
