package db

import (
	"database/sql"
	"fmt"
	"time"

	"diektronics.com/carter/dl/cfg"
	"diektronics.com/carter/dl/types"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type Db struct {
	connectionString string
}

func New(c *cfg.Configuration) *Db {
	return &Db{
		connectionString: fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true&loc=Local",
			c.DbUser, c.DbPassword, c.DbServer, c.DbDatabase),
	}
}

func (d *Db) Add(down *types.Download) error {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	now := time.Now()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("INSERT INTO downloads (name, posthook, created_at, modified_at) VALUES (?, ?, ?, ?)",
		down.Name, down.Posthook, now, now)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	down.ID = id
	down.Status = types.Queued
	down.CreatedAt = now
	down.ModifiedAt = now

	for _, link := range down.Links {
		res, err := tx.Exec("INSERT INTO links (download_id, url, created_at, modified_at) VALUES (?, ?, ?, ?)",
			id, link.Url, now, now)
		if err != nil {
			tx.Rollback()
			return err
		}
		link_id, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			return err
		}
		link.ID = link_id
		link.Status = types.Queued
		link.CreatedAt = now
		link.ModifiedAt = now
	}
	tx.Commit()

	return nil
}

func (d *Db) Get(id int64) (*types.Download, error) {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	down := &types.Download{ID: id}
	var status string
	if err := db.QueryRow("SELECT name, status, error, posthook, created_at, modified_at FROM downloads WHERE id=?", id).Scan(
		&down.Name, &status, &down.Error, &down.Posthook, &down.CreatedAt, &down.ModifiedAt); err != nil {
		return nil, err
	}
	down.Status = types.Status(status)

	rows, err := db.Query("SELECT id, url, status, created_at, modified_at FROM links WHERE download_id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		l := &types.Link{}
		if err := rows.Scan(&l.ID, &l.Url, &status, &l.CreatedAt, &l.ModifiedAt); err != nil {
			return nil, err
		}
		l.Status = types.Status(status)
		down.Links = append(down.Links, l)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return down, nil
}

func (d *Db) Del(down *types.Download) error {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("DELETE FROM links WHERE download_id=?", down.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if n != int64(len(down.Links)) {
		tx.Rollback()
		return fmt.Errorf("unexpected rows affected: %v != %v", n, len(down.Links))
	}

	res, err = tx.Exec("DELETE FROM downloads WHERE id=?", down.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	n, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if n != 1 {
		tx.Rollback()
		return fmt.Errorf("unexpected rows affected: %v != 1", n)
	}

	tx.Commit()
	return nil
}

func (d *Db) Update(down *types.Download) error {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	now := time.Now()
	for _, l := range down.Links {
		_, err = tx.Exec("UPDATE links SET status=?, modified_at=? WHERE id=?",
			string(l.Status), now, l.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	_, err = tx.Exec("UPDATE downloads SET status=?, error=?, modified_at=? WHERE id=?",
		string(down.Status), down.Error, now, down.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
