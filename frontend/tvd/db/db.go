package db

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/diektronics/dl-us/frontend/tvd/show"
	"github.com/diektronics/dl-us/protos/cfg"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type Episode struct {
	Title    string
	Episode  string
	Location string
}

type Db struct {
	connectionString string
}

func New(c *cfg.Config) *Db {
	return &Db{
		connectionString: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
			c.Db.User, c.Db.Password, c.Db.Server, c.Db.Database),
	}
}

func (d *Db) GetMyShows(titles []string) ([]*Episode, error) {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dbQuery := "SELECT name, latest_ep, location FROM series where name IN ("
	vals := []interface{}{}
	for _, t := range titles {
		dbQuery += "?,"
		vals = append(vals, t)
	}
	dbQuery = dbQuery[0 : len(dbQuery)-1]
	dbQuery += ")"
	stmt, err := db.Prepare(dbQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	myShows := []*Episode{}
	for rows.Next() {
		eps := &Episode{}
		err = rows.Scan(&eps.Title, &eps.Episode, &eps.Location)
		if err != nil {
			return nil, err
		}
		myShows = append(myShows, eps)
	}
	return myShows, nil
}

func (d *Db) UpdateMyShows(shows []*show.Show) error {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	var lastErr error
	sort.Slice(shows, func(i, j int) bool {
		return shows[i].Name < shows[j].Name || shows[i].Name == shows[j].Name && shows[i].Eps < shows[j].Eps
	})
	for _, s := range shows {
		dbQuery := fmt.Sprintf("UPDATE series SET latest_ep=%q WHERE name=%q", s.Eps, s.Name)
		_, err = db.Exec(dbQuery)
		if err != nil {
			lastErr = err
			continue
		}
	}

	return lastErr
}
