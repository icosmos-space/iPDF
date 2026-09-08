package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"ipdf/internal/pdf"
)

// App is the Wails binding root.
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct.
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ToolResult is re-exported for Wails bindings.
type ToolResult = pdf.ToolResult

// MergePDFs concatenates PDFs in order.
func (a *App) MergePDFs(paths []string, outPath string) (ToolResult, error) {
	return pdf.Merge(paths, outPath)
}

// SplitPDF splits a PDF. mode: each | range | everyN.
func (a *App) SplitPDF(path, mode, ranges, outDir string) (ToolResult, error) {
	return pdf.Split(path, mode, ranges, outDir)
}

// CompressPDF optimizes a PDF. quality: low | medium | high.
func (a *App) CompressPDF(path, quality, outPath string) (ToolResult, error) {
	return pdf.Compress(path, quality, outPath)
}

// PDFToImages extracts embedded images from a PDF (pdfcpu; not full-page rasterize).
func (a *App) PDFToImages(path, format string, dpi int, outDir string) (ToolResult, error) {
	return pdf.ToImages(path, format, dpi, outDir)
}

// ImagesToPDF builds a PDF from images.
func (a *App) ImagesToPDF(paths []string, outPath string) (ToolResult, error) {
	return pdf.FromImages(paths, outPath)
}

// SelectFiles opens a multi-file dialog. kind: pdf | image | any
func (a *App) SelectFiles(kind string) ([]string, error) {
	return runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title:   "选择文件",
		Filters: dialogFilters(kind),
	})
}

// SelectFile opens a single-file dialog.
func (a *App) SelectFile(kind string) (string, error) {
	return runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:   "选择文件",
		Filters: dialogFilters(kind),
	})
}

// SelectSaveFile opens a save-file dialog.
func (a *App) SelectSaveFile(defaultFilename string) (string, error) {
	return runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "保存文件",
		DefaultFilename: defaultFilename,
		Filters: []runtime.FileFilter{
			{DisplayName: "PDF (*.pdf)", Pattern: "*.pdf"},
		},
	})
}

// SelectDirectory opens a directory picker.
func (a *App) SelectDirectory() (string, error) {
	return runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择输出目录",
	})
}

// OpenInExplorer reveals a file or opens a directory in the OS file manager.
func (a *App) OpenInExplorer(path string) error {
	if path == "" {
		return fmt.Errorf("路径为空")
	}
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	switch goruntime.GOOS {
	case "windows":
		if info.IsDir() {
			return exec.Command("explorer", path).Start()
		}
		return exec.Command("explorer", "/select,", path).Start()
	case "darwin":
		if info.IsDir() {
			return exec.Command("open", path).Start()
		}
		return exec.Command("open", "-R", path).Start()
	default:
		dir := path
		if !info.IsDir() {
			dir = filepath.Dir(path)
		}
		return exec.Command("xdg-open", dir).Start()
	}
}

func dialogFilters(kind string) []runtime.FileFilter {
	switch strings.ToLower(kind) {
	case "pdf":
		return []runtime.FileFilter{{DisplayName: "PDF (*.pdf)", Pattern: "*.pdf"}}
	case "image":
		return []runtime.FileFilter{
			{DisplayName: "图片 (*.jpg;*.jpeg;*.png;*.webp)", Pattern: "*.jpg;*.jpeg;*.png;*.webp"},
		}
	default:
		return []runtime.FileFilter{{DisplayName: "所有文件 (*.*)", Pattern: "*.*"}}
	}
}
