package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type entry struct {
    ID     string  `json:"id"`
    Event string  `json:"event"`
    Message  string `json:"message"`
}

var entries = []entry{
    {ID: "1", Event: "record.created", Message: "A record has been created"},
    {ID: "2", Event: "record.modified", Message: "A recod has been modified"},
    {ID: "3", Event: "record.archived", Message: "A record has been archived"},
}


func listEntries(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, entries)
}

func postEntry(c *gin.Context) {
	var newEntry entry

	if err := c.BindJSON(&newEntry); err != nil {
        return
    }

	entries = append(entries, newEntry)
	c.IndentedJSON(http.StatusCreated, newEntry)
}

func main() {
    router := gin.Default()
    router.GET("/entries", listEntries)
	router.POST("/entries", postEntry)

    router.Run("localhost:8080")
}