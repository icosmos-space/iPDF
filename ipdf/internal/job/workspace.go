package job

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Workspace is a per-job temporary directory under os.TempDir()/ipdf/<id>.
type Workspace struct {
	ID   string
	Root string
}

// New creates a fresh workspace directory.
func New() (*Workspace, error) {
	id := uuid.NewString()
	root := filepath.Join(os.TempDir(), "ipdf", id)
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, fmt.Errorf("create workspace: %w", err)
	}
	return &Workspace{ID: id, Root: root}, nil
}

// Path joins name under the workspace root.
func (w *Workspace) Path(name string) string {
	return filepath.Join(w.Root, name)
}

// Cleanup removes the workspace directory.
func (w *Workspace) Cleanup() {
	if w == nil || w.Root == "" {
		return
	}
	_ = os.RemoveAll(w.Root)
}

// CopyFile copies src to dst, creating parent directories as needed.
func CopyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

// MoveOrCopy moves src to dst, falling back to copy+remove across volumes.
func MoveOrCopy(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	if err := os.Rename(src, dst); err == nil {
		return nil
	}
	if err := CopyFile(src, dst); err != nil {
		return err
	}
	_ = os.Remove(src)
	return nil
}
