package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"ipdf/internal/job"
)

// Compress optimizes a PDF. quality: low | medium | high (hints for pdfcpu optimize).
func Compress(path, quality, outPath string) (ToolResult, error) {
	if err := validatePDFInput(path); err != nil {
		return ToolResult{}, err
	}
	if outPath == "" {
		return ToolResult{}, fmt.Errorf("请指定输出路径")
	}

	ws, err := job.New()
	if err != nil {
		return ToolResult{}, err
	}
	defer ws.Cleanup()

	tmpOut := ws.Path("compressed.pdf")
	c := conf()
	switch quality {
	case "low", "medium", "high", "":
		// pdfcpu OptimizeFile removes unused objects; quality is a UX hint for now.
	default:
		return ToolResult{}, fmt.Errorf("未知压缩等级: %s（支持 low / medium / high）", quality)
	}

	if err := api.OptimizeFile(path, tmpOut, c); err != nil {
		return ToolResult{}, fmt.Errorf("压缩失败: %w", err)
	}
	if err := job.MoveOrCopy(tmpOut, outPath); err != nil {
		return ToolResult{}, fmt.Errorf("写入输出文件失败: %w", err)
	}

	msg := fmt.Sprintf("压缩完成（%s）→ %s", qualityOrDefault(quality), filepath.Base(outPath))
	if inStat, err1 := fileSize(path); err1 == nil {
		if outStat, err2 := fileSize(outPath); err2 == nil && inStat > 0 {
			pct := float64(outStat) / float64(inStat) * 100
			msg = fmt.Sprintf("压缩完成：%.1f KB → %.1f KB（约 %.0f%%）",
				float64(inStat)/1024, float64(outStat)/1024, pct)
		}
	}

	return ToolResult{
		OutputPaths: []string{outPath},
		Message:     msg,
	}, nil
}

func qualityOrDefault(q string) string {
	if q == "" {
		return "medium"
	}
	return q
}

func fileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
