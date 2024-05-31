package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ayushkumarone/UrlShortener/requests"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default() // Defining the router using gin framework.

	// Accessing .env file
	godotenv.Load(".env")
	// Verify the access
	// fmt.Println(os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("Address"), os.Getenv("DB_Name"))

	// ----------------	 START : Defining the configurations  ----------------

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("Address"),
		DBName:               os.Getenv("DB_Name"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//  ----------------  END : Defining the configurations  ----------------

	//  ----------------  START : Verify connection  ----------------

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		return
	}
	fmt.Println("Successfully connected to the MariaDB database!")

	//  ----------------  END : Verify connection  ----------------

	//  ----------------  START : Automated unused URL removal  ----------------

	//  ----------------  END : Automated unused URL removal  ----------------

	//  ----------------  START : Routes defined here  ----------------

	router.GET("/link/:id", func(c *gin.Context) {
		requests.GetLinkByID(c, db)
	})

	router.POST("/analytics/:id", func(c *gin.Context) {
		requests.CountByID(c, db)
	})

	//  ----------------  END : Routes defined here  ----------------

	router.Run("localhost:8080") // Running router on localhost port 8080
}
