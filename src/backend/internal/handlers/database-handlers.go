package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Database struct {
	Name string `json:"name"`
}

// CreateDatabase godoc
// @Summary Create a database
// @Description Create a new database with the given name
// @Param database body Database true "Database name"
// @Produce application/json
// @Tags Database
// @Success         200 {object} map[string]string
// @Failure         400 {object} map[string]string
// @Router /database/ [post]
func CreateDatabase(c *gin.Context) {
	var db Database
	if err := c.ShouldBindJSON(&db); err != nil || db.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, database name is required"})
		return
	}
	c.JSON(http.StatusCreated, db)
}

// GetAllDatabases godoc
// @Summary         Get all databases
// @Description     Get all databases
// @Produce         application/json
// @Tags            Database
// @Success         200 {array} Database
// @Router          /database/getall [get]
func GetAllDatabases(c *gin.Context) {
}
