package database

import (
	"github.com/doug-martin/goqu/v9"
	// import the dialect.
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

func GetGoquDialect(connection string) goqu.DialectWrapper {
	var qu goqu.DialectWrapper

	switch connection {
	case MySQL:
		qu = goqu.Dialect(MySQL)
	case SQLite:
		qu = goqu.Dialect(SQLite)
	case PostgreSQL:
		qu = goqu.Dialect(PostgreSQL)
	}

	return qu
}
