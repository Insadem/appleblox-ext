package backup

import (
	"io"
	"os"
)

type Backup struct {
	dir string
}

func (b Backup) AppPath() string {
	return b.dir + "/Roblox.app"
}

// Close releases all associated resources with the backup.
func (b Backup) Close() error {
	return os.RemoveAll(b.dir)
}

var _ io.Closer = (*Backup)(nil)
