// +build integration

package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-postgres-cicd/database"
)

func TestGetProducts_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &database.Product{Name: productName, Price: productPrice}
	insertErr := database.CreateProduct(db, product, ctx)
	require.NoError(t, insertErr)

	products, err := database.GetProducts(db, 0, 10, ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 1)
	assert.GreaterOrEqual(t, products[0].ID, 0)
	assert.Equal(t, productName, products[0].Name)
	assert.Equal(t, productPrice, products[0].Price)

	database.DeleteProducts(db, ctx)
}

func TestGetProduct_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	sourceProd := &database.Product{Name: productName, Price: productPrice}
	insertErr := database.CreateProduct(db, sourceProd, ctx)
	require.NoError(t, insertErr)

	targetProd := &database.Product{ID: sourceProd.ID}
	err := database.GetProduct(db, targetProd, ctx)
	assert.NoError(t, err)
	assert.Equal(t, sourceProd.ID, targetProd.ID)
	assert.Equal(t, sourceProd.Name, targetProd.Name)
	assert.Equal(t, sourceProd.Price, targetProd.Price)

	database.DeleteProducts(db, ctx)
}

func TestCreateProduct_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &database.Product{Name: productName, Price: productPrice}
	err := database.CreateProduct(db, product, ctx)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, product.ID, 0)
	assert.Equal(t, productName, product.Name)
	assert.Equal(t, productPrice, product.Price)

	database.DeleteProducts(db, ctx)
}

func TestUpdateProduct_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	insert := &database.Product{Name: productName, Price: productPrice}
	insertErr := database.CreateProduct(db, insert, ctx)
	require.NoError(t, insertErr)

	update := &database.Product{ID: insert.ID, Name: productNewName, Price: productNewPrice}
	err := database.UpdateProduct(db, update, ctx)
	assert.NoError(t, err)
	assert.Equal(t, insert.ID, update.ID)
	assert.Equal(t, productNewName, update.Name)
	assert.Equal(t, productNewPrice, update.Price)
	assert.NotEqual(t, insert.Name, update.Name)
	assert.NotEqual(t, insert.Price, update.Price)

	database.DeleteProducts(db, ctx)
}

func TestDeleteProduct_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &database.Product{Name: productName, Price: productPrice}
	insertErr := database.CreateProduct(db, product, ctx)
	require.NoError(t, insertErr)

	getErr := database.GetProduct(db, product, ctx)
	require.NoError(t, getErr)

	err := database.DeleteProduct(db, product.ID, ctx)
	assert.NoError(t, err)

	database.DeleteProducts(db, ctx)
}

func TestDeleteProducts_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &database.Product{Name: "one", Price: 1.10}
	insertErr := database.CreateProduct(db, product, ctx)
	require.NoError(t, insertErr)

	product2 := &database.Product{Name: "two", Price: 2.20}
	insert2Err := database.CreateProduct(db, product2, ctx)
	require.NoError(t, insert2Err)

	product3 := &database.Product{Name: "three", Price: 3.30}
	insert3Err := database.CreateProduct(db, product3, ctx)
	require.NoError(t, insert3Err)

	productsBefore, getErrBefore := database.GetProducts(db, 0, 10, ctx)
	require.NoError(t, getErrBefore)
	require.Len(t, productsBefore, 3)

	err := database.DeleteProducts(db, ctx)
	assert.NoError(t, err)

	productsAfter, getErrAfter := database.GetProducts(db, 0, 10, ctx)
	assert.NoError(t, getErrAfter)
	assert.Len(t, productsAfter, 0)

	database.DeleteProducts(db, ctx)
}
