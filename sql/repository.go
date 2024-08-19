package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func NewManager(path string, ctx context.Context) (*Manager, error) {
	manager := &Manager{
		path:         path,
		removeOnExit: false,
		ctx:          ctx,
	}

	db, err := sql.Open("libsql", fmt.Sprintf("file:%s", path))

	if err != nil {
		return nil, err
	}

	manager.db = db

	return manager, nil
}

type Manager struct {
	db           *sql.DB
	path         string
	removeOnExit bool
	ctx          context.Context
}

func (this *Manager) Cleanup() error {
	closeErr := this.db.Close()
	var removeErr error

	if this.removeOnExit {
		removeErr = os.Remove(this.path)
	}

	return errors.Join(closeErr, removeErr)
}

func (this *Manager) EnsureTables(schema DatabaseSchema) error {
	ddl := schema.Generate()

	_, err := this.db.ExecContext(this.ctx, ddl)

	if err != nil {
		return nil
	}

	return nil
}
