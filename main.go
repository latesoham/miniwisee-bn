package main

import ("database/sql"
		"fmt"
		"log"

		"github.com/gin-gonic/gin"
		"github.com/gin-contrib/cors"
		_ "github.com/go-sql-driver/mysql"
	)
type Ticket struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Tag      string `json:"tag"`
	Status   string `json:"status"`
}

var tickets []Ticket

var db *sql.DB

func main() {

var err error

db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/miniwise")

if err != nil {
	log.Fatal(err)
}

err = db.Ping()

if err != nil {
	log.Fatal(err)
}

fmt.Println("✅ Connected to MySQL")

	router := gin.Default()

	router.Use(cors.Default())

router.GET("/", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "MiniWise Backend Running 🚀",
	})
})

router.GET("/tickets", func(c *gin.Context) {

	tickets := []Ticket{}

	rows, err := db.Query("SELECT id, title, priority, tag_name, status FROM tickets")

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {

		var ticket Ticket

		err := rows.Scan(
			&ticket.ID,
			&ticket.Title,
			&ticket.Priority,
			&ticket.Tag,
			&ticket.Status,
		)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		tickets = append(tickets, ticket)
	}

	c.JSON(200, tickets)
})

router.POST("/tickets", func(c *gin.Context) {

	var newTicket Ticket

	if err := c.BindJSON(&newTicket); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	query := `
	INSERT INTO tickets (title, priority, tag_name, status)
	VALUES (?, ?, ?, ?)
	`

	result, err := db.Exec(
		query,
		newTicket.Title,
		newTicket.Priority,
		newTicket.Tag,
		newTicket.Status,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := result.LastInsertId()

	newTicket.ID = int(id)

	c.JSON(201, gin.H{
		"message": "Ticket created successfully",
		"ticket":  newTicket,
	})
})
router.DELETE("/tickets/:id", func(c *gin.Context) {

	id := c.Param("id")

	query := "DELETE FROM tickets WHERE id = ?"

	result, err := db.Exec(query, id)
	rowsAffected, _ := result.RowsAffected()

	fmt.Println("Deleted rows:", rowsAffected)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Ticket deleted successfully",
	})
})
	router.Run(":8080")
}