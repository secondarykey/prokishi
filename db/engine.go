package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

const EnginesColumns = "id,path,created_date,updated_date"
const EnginesSelect = "SELECT id,path,DATETIME(created_date),DATETIME(updated_date) FROM engines"

type Engine struct {
	ID      string
	Path    string
	Created time.Time
	Updated time.Time
}

func createEngine(row scanner) (*Engine, error) {
	var e Engine
	err := row.Scan(&e.ID, &e.Path, &e.Created, &e.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, xerrors.Errorf("Scan() error: %w", err)
	}
	return &e, nil
}

func SelectEngine(ctx context.Context, id string) (*Engine, error) {
	s := EnginesSelect + " WHERE id = ?"
	row, err := getRow(ctx, s, id)
	if err != nil {
		return nil, xerrors.Errorf("getRow(engine) error: %w", err)
	}
	if row != nil {
		return createEngine(row)
	}
	return nil, nil
}

func FindEngines(ctx context.Context) ([]*Engine, error) {
	s := EnginesSelect + " ORDER BY updated_date DESC"
	rows, err := getRows(ctx, s)
	if err != nil {
		return nil, xerrors.Errorf("getRows(engine) error: %w", err)
	}
	engines := make([]*Engine, 0)
	for rows.Next() {
		e, err := createEngine(rows)
		if err != nil {
			return nil, xerrors.Errorf("createEngine() error: %w", err)
		} else if e == nil {
			break
		}
		engines = append(engines, e)
	}
	return engines, nil
}

func InsertEngine(id string, path string) error {
	now := time.Now()
	s := fmt.Sprintf("INSERT INTO engines (%s) VALUES (?,?,?,?)", EnginesColumns)
	err := run(s, id, path, now, now)
	if err != nil {
		return xerrors.Errorf("run() error: %w", err)
	}
	return nil
}

func DeleteEngine(id string) error {
	s := "DELETE FROM engines WHERE id = ?"
	err := run(s, id)
	if err != nil {
		return xerrors.Errorf("run() error: %w", err)
	}
	return nil
}
