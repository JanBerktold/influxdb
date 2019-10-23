package influxdb

import (
	"context"
	"io"
)

type BackupService interface {
	CreateBackup(context.Context) (int, []string, error)
	FetchBackupFile(ctx context.Context, backupID int, backupFile string, w io.Writer) error
	InternalBackupPath(backupID int) string
}

type KVBackupService interface {
	Backup(ctx context.Context, w io.Writer) error
}
