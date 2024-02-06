package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/m1guelsb/eshop-goapi/internal/database"
	"github.com/m1guelsb/eshop-goapi/internal/service"
	"github.com/m1guelsb/eshop-goapi/internal/webserver"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/eshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Get("/category/{id}", webCategoryHandler.GetCategoryByID)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/{id}", webProductHandler.GetProductByID)
	c.Get("/product/category/{id}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server running on port 3001")
	http.ListenAndServe(":3001", c)
}
