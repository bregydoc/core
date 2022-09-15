package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minskylab/core/ent"
	"github.com/pkg/errors"
	"github.com/xo/dburl"
)

func openEntClient(dbURL string) (*ent.Client, error) {
	db, err := dburl.Open(dbURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var driver *entsql.Driver = nil

	if strings.HasPrefix(dbURL, "postgres") {
		driver = entsql.OpenDB(dialect.Postgres, db)
	} else if strings.HasPrefix(dbURL, "sqlite") {
		driver = entsql.OpenDB(dialect.SQLite, db)
	}

	if driver == nil {
		return nil, fmt.Errorf("invalid sql dialect. '%s' url not supported", dbURL)
	}

	return ent.NewClient(ent.Driver(driver)), nil
}

// OpenDBClient create a new Ent Client with Postgres or Sqlite Connection
func OpenDBClient(dbURL string) (*ent.Client, error) {
	return openEntClient(dbURL)
}
