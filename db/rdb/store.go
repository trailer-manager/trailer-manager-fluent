package db

import (
	"SiverPineValley/trailer-manager/config"
	"context"
	"database/sql"
	"fmt"
	"log"
)

var RDB *sql.DB

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute db queries and transactions.
type SQLStore struct {
	db *sql.DB
	*Queries
}

func getDbSource(conf config.Rdb) (source string) {
	switch conf.Driver {
	case "postgres":
		source = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	}
	return
}

func InitRdb() (err error) {
	conf := config.GetConfig().Rdb
	source := getDbSource(conf)
	conn, err := sql.Open(conf.Driver, source)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	RDB = conn
	return
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction.
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}