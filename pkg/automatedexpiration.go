package pkg

import (
	"database/sql"
	"fmt"
	"time"
)

func RemoveLinks(db *sql.DB) {
	time.Sleep(180)
	deletingUnused := "DELETE FROM shorturls WHERE TIMESTAMPDIFF(SECOND, LastHit, CURRENT_TIMESTAMP) > 180;"
	_, err2 := db.Exec(deletingUnused) // Performing the query.
	if err2 != nil {
		// fmt.Println("Error at performing deletingUnused: ", err2)
		fmt.Println("Error deleting the unused shortened links.")
		return
	}
	RemoveLinks(db)
}
