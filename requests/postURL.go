package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/UrlShortener/pkg/links"

	"github.com/gin-gonic/gin"
)

func PostURL(c *gin.Context, db *sql.DB) {
	var newLink links.Link

	// Extracting the JSON
	if err := c.BindJSON(&newLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	// Checking if that id exists previously
	query := fmt.Sprintf("SELECT Short FROM shorturls where Short='%v';", newLink.Short) // Query for checking

	// Performing the query
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var duplicate int
	for rows.Next() {
		var short string
		if err := rows.Scan(&short); err != nil {
			fmt.Print(err)
			return
		}
		duplicate++
	}
	// duplicate > 0 means same id exists
	if duplicate > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Short url already exist"})
		return
	}

	// performing the insertion
	queryInsert := fmt.Sprintf("INSERT INTO shorturls (Short, Url) VALUES ('%v', '%v');", newLink.Short, newLink.Url) // query for insertion
	_, err2 := db.Exec(queryInsert)
	// Handling the error
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	}

}
