package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-sb-pekan3/controllers"
	"quiz-sb-pekan3/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "1234"
// 	dbname   = "practice"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                // Mengizinkan semua origin, Anda bisa menambahkan domain frontend spesifik
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},     // Metode HTTP yang diperbolehkan
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Header yang diperbolehkan
		AllowCredentials: true,                                         // Jika Anda membutuhkan kredensial (misalnya cookies)
	}))

	api := router.Group("/api")

	api.Use(controllers.BasicAuthMiddleware(DB)) // Menambahkan BasicAuth di grup ini

	{
		// Semua rute di bawah /api akan memerlukan Basic Authentication
		api.GET("/categories", controllers.GetAllCategory)
		api.POST("/categories", controllers.InsertCategory)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.GET("/categories/:id", controllers.GetDetailCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)
		api.GET("/categories/:id/books", controllers.GetDetailBookbyCategory)

		api.POST("/books", controllers.InsertBook)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.GET("/books/:id", controllers.GetDetailBook)
		api.DELETE("/books/:id", controllers.DeleteBook)
	}
	router.GET("/books", controllers.GetAllBook)
	router.Run(":" + os.Getenv("PORT"))
	//router.Run(":8080")
}
