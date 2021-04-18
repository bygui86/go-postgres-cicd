package database_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

const (
	createTableQuery   = "CREATE TABLE IF NOT EXISTS products"
	getProductsQuery   = "SELECT id,name,price FROM products"
	getProductQuery    = "SELECT name,price FROM products"
	createProductQuery = "INSERT INTO products"
	updateProductQuery = "UPDATE products"
	deleteProductQuery = "DELETE FROM products"
)

/*
	By default, sqlmock is preserving backward compatibility and default query matcher is sqlmock.QueryMatcherRegexp
	which uses expected SQL string as a regular expression to match incoming query string.
*/
func NewRegexpMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	return db, mock
}

/*
	sqlmock.QueryMatcherEqual which will do a full case sensitive match.
*/
func NewEqualMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	return db, mock
}
