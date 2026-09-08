package pdf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"ipdf/internal/job"
)

// Split splits a PDF.
// mode: each | range | everyN
// ranges: for range e.g. "1-3,5"; for everyN e.g. "2"
func Split(path, mode, ranges, outDir string) (ToolResult, error) {
	if err := validatePDFInput(path); err != nil {
		return ToolResult{}, err
	}
	if outDir == "" {
		return ToolResult{}, fmt.Errorf("请指定输出目录")
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return ToolResult{}, fmt.Errorf("创建输出目录失败: %w", err)
	}

	ws, err := job.New()
	if err != nil {
		return ToolResult{}, err
	}
	defer ws.Cleanup()

	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	tmpDir := ws.Path("split")
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		return ToolResult{}, err
	}

	switch mode {
	case "each":
		if err := api.SplitFile(path, tmpDir, 1, conf()); err != nil {
			return ToolResult{}, fmt.Errorf("按页分割失败: %w", err)
		}
	case "everyN":
		n, err := strconv.Atoi(strings.TrimSpace(ranges))
		if err != nil || n < 1 {
			return ToolResult{}, fmt.Errorf("everyN 需要正整数，当前: %q", ranges)
		}
		if err := api.SplitFile(path, tmpDir, n, conf()); err != nil {
			return ToolResult{}, fmt.Errorf("按 N 页分割失败: %w", err)
		}
	case "range":
		sel := strings.TrimSpace(ranges)
		if sel == "" {
			return ToolResult{}, fmt.Errorf("请指定页码范围，例如 1-3,5")
		}
		tmpOut := filepath.Join(tmpDir, base+"_extract.pdf")
		if err := api.TrimFile(path, tmpOut, []string{sel}, conf()); err != nil {
			return ToolResult{}, fmt.Errorf("按范围提取失败: %w", err)
		}
	default:
		return ToolResult{}, fmt.Errorf("未知分割模式: %s（支持 each / range / everyN）", mode)
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		return ToolResult{}, err
	}
	var outs []string
	for _, e := range entries {
		if e.IsDir() || !strings.EqualFold(filepath.Ext(e.Name()), ".pdf") {
			continue
		}
		src := filepath.Join(tmpDir, e.Name())
		dst := filepath.Join(outDir, e.Name())
		if err := job.MoveOrCopy(src, dst); err != nil {
			return ToolResult{}, err
		}
		outs = append(outs, dst)
	}
	if len(outs) == 0 {
		return ToolResult{}, fmt.Errorf("未生成任何输出文件")
	}

	return ToolResult{
		OutputPaths: outs,
		Message:     fmt.Sprintf("已生成 %d 个文件到 %s", len(outs), outDir),
	}, nil
}
