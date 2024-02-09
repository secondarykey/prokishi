package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

const CodesColumns = "code,created_date,updated_date"
const CodesSelect = "SELECT code,DATETIME(created_date),DATETIME(updated_date) FROM codes"

type Code struct {
	Code    string
	Created time.Time
	Updated time.Time
}

func createCode(row scanner) (*Code, error) {
	var c Code
	err := row.Scan(&c.Code, &c.Created, &c.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, xerrors.Errorf("Scan() error: %w", err)
	}
	return &c, nil
}

func SelectCode(ctx context.Context, code string) (*Code, error) {
	s := CodesSelect + " WHERE code = ?"
	row, err := getRow(ctx, s, code)
	if err != nil {
		return nil, xerrors.Errorf("getRow(code) error: %w", err)
	}
	if row != nil {
		return createCode(row)
	}
	return nil, nil
}

func FindCodes(ctx context.Context) ([]*Code, error) {
	s := CodesSelect + " ORDER BY updated_date DESC"
	rows, err := getRows(ctx, s)
	if err != nil {
		return nil, xerrors.Errorf("getRows(code) error: %w", err)
	}
	codes := make([]*Code, 0)
	for rows.Next() {
		c, err := createCode(rows)
		if err != nil {
			return nil, xerrors.Errorf("createCode() error: %w", err)
		} else if c == nil {
			break
		}
		codes = append(codes, c)
	}
	return codes, nil
}

func InsertCode(code string) error {
	now := time.Now()
	zero := time.Time{}
	s := fmt.Sprintf("INSERT INTO codes (%s) VALUES (?,?,?)", CodesColumns)
	err := run(s, code, now, zero)
	if err != nil {
		return xerrors.Errorf("run() error: %w", err)
	}
	return nil
}

func UpdateCode(code string) error {
	now := time.Now()
	s := "UPDATE codes SET updated_date = ? WHERE code = ?"
	err := run(s, now, code)
	if err != nil {
		return xerrors.Errorf("run() error: %w", err)
	}
	return nil
}

func DeleteCode(code string) error {
	s := "DELETE FROM codes WHERE code = ?"
	err := run(s, code)
	if err != nil {
		return xerrors.Errorf("run() error: %w", err)
	}
	return nil
}
