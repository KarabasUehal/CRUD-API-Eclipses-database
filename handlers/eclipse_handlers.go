package handlers

import (
	"net/http"
	"strconv"

	"golang-gin/storage"

	"github.com/gin-gonic/gin"
)

func GetAllEclipses(c *gin.Context) {
	eclipses := storage.GetAllEclipses()

	c.JSON(http.StatusOK, eclipses)
}

func GetEclipseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid eclipse ID"})
		return
	}

	eclipse := storage.GetEclipseByID(id)

	if eclipse == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Eclipse not found"})
		return
	}
	c.JSON(http.StatusOK, eclipse)
}

func AddEclipse(c *gin.Context) {
	var input struct {
		Description string  `json:"description"`
		Duration    float64 `json:"duration"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	eclipse := storage.AddEclipse(input.Description, input.Duration)

	c.JSON(http.StatusCreated, eclipse)
}

func UpdateEclipseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid eclipse ID",
		})
		return
	}

	var input struct {
		Description string  `json:"description"`
		Duration    float64 `json:"duration"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	eclipse := storage.UpdateEclipseByID(id, input.Description, input.Duration)

	if eclipse == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Eclipse not found"})
		return
	}
	c.JSON(http.StatusOK, eclipse)
}

func DeleteEclipseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaled eclipse ID",
		})
		return
	}
	if success := storage.DeleteEclipseByID(id); !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Eclipse not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
