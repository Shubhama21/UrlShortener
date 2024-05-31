package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CountByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	queryHandle := fmt.Sprintf("SELECT Short FROM shorturls where Short='%v';", id)

	rows2, err2 := db.Query(queryHandle)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows2.Close()

	unique := 0
	for rows2.Next() {
		var short string
		if err2 := rows2.Scan(&short); err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		unique++
	}
	if unique == 0 || unique > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "This id does not exist."})
	}

	query := fmt.Sprintf("Select HitCount from shorturls where Short='%v';", id)

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Hit_Count : ": count})
	}

}
