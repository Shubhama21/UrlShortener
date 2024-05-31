package requests

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLinkByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id") // Extracting the id from the path for finding in the database.

	query := fmt.Sprintf("SELECT Url FROM shorturls where Short='%v';", id) // Framing the SQL query for finding the URL.
	// fmt.Println("My query is: ",query)

	rows, err1 := db.Query(query) // Performing the query.
	if err1 != nil {
		// fmt.Println("Error at performing query: ", err1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var url string
	if rows.Next() {
		if err := rows.Scan(&url); err != nil {
			// fmt.Println("Error at scanning rows: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else {
		// No rows are returned, Short not in database
		c.JSON(http.StatusNotFound, gin.H{"message": "link not found"})
		return
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	c.Redirect(http.StatusSeeOther, url)
}
