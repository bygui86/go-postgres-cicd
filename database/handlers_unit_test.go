// +build !integration

package database_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-postgres-cicd/database"
	"github.com/bygui86/go-postgres-cicd/logging"
)

func TestGetProducts_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(productId, productName, productPrice).
		AddRow(productId2, productName2, productPrice2)

	mock.ExpectQuery(getProductsQuery).
		WillReturnRows(rows)

	products, err := database.GetProducts(db, 0, 10, context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 2)
	assert.Equal(t, productId, products[0].ID)
	assert.Equal(t, productName, products[0].Name)
	assert.Equal(t, productPrice, products[0].Price)
	assert.Equal(t, productId2, products[1].ID)
	assert.Equal(t, productName2, products[1].Name)
	assert.Equal(t, productPrice2, products[1].Price)
}

func TestGetProducts_Unit_Fail_Query(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectQuery(getProductsQuery).
		WillReturnError(fmt.Errorf("error"))

	products, err := database.GetProducts(db, 0, 10, context.Background())

	assert.Error(t, err)
	assert.Equal(t, 0, len(products))
}

// see https://github.com/DATA-DOG/go-sqlmock/issues/47
func TestGetProducts_Unit_Fail_Scan(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(productId, productName, productPrice).
		AddRow(productId2, productName2, productPrice2).
		AddRow(nil, "sample-3", 44.44).RowError(3, fmt.Errorf("row-error"))

	mock.ExpectQuery(getProductsQuery).
		WillReturnRows(rows)

	products, err := database.GetProducts(db, 0, 10, context.Background())

	assert.Error(t, err)
	assert.Equal(t, 0, len(products))
}

func TestGetProduct_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "price"}).
		AddRow(productName, productPrice)

	mock.ExpectQuery(getProductQuery).
		WithArgs(productId).
		WillReturnRows(rows)

	product := &database.Product{ID: productId}
	err := database.GetProduct(db, product, context.Background())

	assert.NoError(t, err)
	assert.Equal(t, productId, product.ID)
	assert.Equal(t, productName, product.Name)
	assert.Equal(t, productPrice, product.Price)
}

func TestGetProduct_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectQuery(getProductQuery).
		WithArgs(productId).
		WillReturnError(fmt.Errorf("error"))

	product := &database.Product{ID: productId}
	err := database.GetProduct(db, product, context.Background())

	assert.Error(t, err)
	assert.Equal(t, productId, product.ID)
	assert.Equal(t, "", product.Name)
	assert.Equal(t, 0.0, product.Price)
}

func TestCreateProduct_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(productId)

	mock.ExpectQuery(createProductQuery).
		WithArgs(productName, productPrice).
		WillReturnRows(rows)
	mock.ExpectCommit()

	product := &database.Product{Name: productName, Price: productPrice}
	err := database.CreateProduct(db, product, context.Background())

	assert.NoError(t, err)
	assert.Equal(t, productId, product.ID)
	assert.Equal(t, productName, product.Name)
	assert.Equal(t, productPrice, product.Price)
}

func TestCreateProduct_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectQuery(createProductQuery).
		WithArgs(productName, productPrice).
		WillReturnError(fmt.Errorf("error"))
	mock.ExpectRollback()

	product := &database.Product{Name: productName, Price: productPrice}
	err := database.CreateProduct(db, product, context.Background())

	assert.Error(t, err)
}

func TestUpdateProduct_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(updateProductQuery).
		WithArgs(productName, productPrice, productId).
		WillReturnResult(sqlmock.NewResult(productId, 1))
	mock.ExpectCommit()

	product := &database.Product{ID: productId, Name: productName, Price: productPrice}
	err := database.UpdateProduct(db, product, context.Background())

	assert.NoError(t, err)
}

func TestUpdateProduct_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(updateProductQuery).
		WithArgs(productName, productPrice, productId).
		WillReturnError(fmt.Errorf("error"))
	mock.ExpectRollback()

	product := &database.Product{ID: productId, Name: productName, Price: productPrice}
	err := database.UpdateProduct(db, product, context.Background())

	assert.Error(t, err)
}

func TestDeleteProduct_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(deleteProductQuery).
		WithArgs(productId).
		WillReturnResult(sqlmock.NewResult(productId, 1))
	mock.ExpectCommit()

	err := database.DeleteProduct(db, productId, context.Background())

	assert.NoError(t, err)
}

func TestDeleteProduct_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(deleteProductQuery).
		WithArgs(productId).
		WillReturnError(fmt.Errorf("error"))
	mock.ExpectRollback()

	err := database.DeleteProduct(db, productId, context.Background())

	assert.Error(t, err)
}

func TestDeleteProducts_Unit_Success(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(deleteProductQuery).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	err := database.DeleteProducts(db, context.Background())

	assert.NoError(t, err)
}

func TestDeleteProducts_Unit_Fail(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	db, mock := NewRegexpMock(t)
	defer db.Close()

	mock.ExpectExec(deleteProductQuery).
		WithArgs(productId).
		WillReturnError(fmt.Errorf("error"))
	mock.ExpectRollback()

	err := database.DeleteProducts(db, context.Background())

	assert.Error(t, err)
}
