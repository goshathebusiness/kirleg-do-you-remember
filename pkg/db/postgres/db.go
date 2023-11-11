package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
	isolationLevel sql.IsolationLevel
}

var IsolationLevels = map[string]sql.IsolationLevel{
	"Default":        sql.LevelDefault,
	"ReadUncommitte": sql.LevelReadUncommitted,
	"ReadCommitted":  sql.LevelReadCommitted,
	"WriteCommitted": sql.LevelWriteCommitted,
	"RepeatableRead": sql.LevelRepeatableRead,
	"Snapshot":       sql.LevelSnapshot,
	"Serializable":   sql.LevelSerializable,
	"Linearizable":   sql.LevelLinearizable,
}

func NewDB(connString string, isolationLevel string) (*DB, error) {
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		return nil, err
	}

	isolation, exist := IsolationLevels[isolationLevel]
	if !exist {
		return nil, fmt.Errorf("invalid isolation level")
	}

	return &DB{
		DB:             db,
		isolationLevel: isolation,
	}, nil
}

func (db *DB) Begintx(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: db.isolationLevel,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
