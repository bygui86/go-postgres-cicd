// +build !integration

package database_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-postgres-cicd/database"
	"github.com/bygui86/go-postgres-cicd/logging"
)

func TestInitDb_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(createTableQuery).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	err := database.InitDb(db)

	assert.NoError(t, err)
}

func TestInitDb_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(getProductsQuery).
		WillReturnError(fmt.Errorf("error"))

	err := database.InitDb(db)

	assert.Error(t, err)
}
