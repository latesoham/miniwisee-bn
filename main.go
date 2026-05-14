package main

import ("fmt"
		"github.com/gin-gonic/gin"
	)
type Ticket struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Tag      string `json:"tag"`
	Status   string `json:"status"`
}

var tickets = []Ticket{
	{
		ID:       1,
		Title:    "Learn React",
		Priority: "High",
		Tag:      "React",
		Status:   "Todo",
	},
	{
		ID:       2,
		Title:    "Build Backend",
		Priority: "Medium",
		Tag:      "Backend",
		Status:   "In Progress",
	},
}

func main() {
	router := gin.Default()

router.GET("/", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "MiniWise Backend Running 🚀",
	})
})

router.GET("/tickets", func(c *gin.Context) {
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

	tickets = append(tickets, newTicket)

	c.JSON(201, gin.H{
		"message": "Ticket created successfully",
		"ticket": newTicket,
	})
})
router.DELETE("/tickets/:id", func(c *gin.Context) {

	id := c.Param("id")

	var updatedTickets []Ticket

	for _, ticket := range tickets {

		if fmt.Sprint(ticket.ID) != id {
			updatedTickets = append(updatedTickets, ticket)
		}
	}

	tickets = updatedTickets

	c.JSON(200, gin.H{
		"message": "Ticket deleted successfully",
	})
})
	router.Run(":8080")
}