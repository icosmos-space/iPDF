package pdf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func conf() *model.Configuration {
	c := model.NewDefaultConfiguration()
	c.ValidationMode = model.ValidationRelaxed
	return c
}

func ensurePDFExt(path string) error {
	if strings.ToLower(filepath.Ext(path)) != ".pdf" {
		return fmt.Errorf("需要 PDF 文件: %s", path)
	}
	return nil
}

func fileReadable(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("无法访问文件: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("路径是目录而非文件: %s", path)
	}
	return nil
}

func validatePDFInput(path string) error {
	if err := ensurePDFExt(path); err != nil {
		return err
	}
	if err := fileReadable(path); err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("无法打开文件: %w", err)
	}
	defer f.Close()

	info, err := api.PDFInfo(f, path, nil, false, conf())
	if err != nil {
		if errors.Is(err, pdfcpu.ErrEncrypted) || strings.Contains(strings.ToLower(err.Error()), "encrypted") {
			return fmt.Errorf("文件已加密，第一期暂不支持解密: %s", filepath.Base(path))
		}
		// Allow proceed; operation itself will surface a clearer error.
		return nil
	}
	if info != nil && info.Encrypted {
		return fmt.Errorf("文件已加密，第一期暂不支持解密: %s", filepath.Base(path))
	}
	return nil
}
