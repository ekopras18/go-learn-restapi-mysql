package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"url":   "https://ekopras18.com/api/v1",
		"data":  "RESTful API basics: Golang (gin, gorm) x Mysql",
		"alive": true})
}

func ResponseIndex(err error, data interface{}, c *gin.Context) {

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  false,
			"message": "Failed to retrieve data from the database " + err.Error(),
			"data":    []any{},
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "success get data",
		"data":    data,
	})
}

func ResponseValidation(err error, c *gin.Context) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to bind JSON " + err.Error(),
		})
		return
	}

}

func ResponseCreate(err error, data interface{}, c *gin.Context) {

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create data " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success create data",
	})
}

func ResponseUpdate(rowsAffected int64, id string, data interface{}, c *gin.Context) {

	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success updated data",
	})
}

func ResponseShow(err error, data interface{}, c *gin.Context) {
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  true,
				"message": "Record not found!",
				"data":    []any{},
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Failed to retrieve data from the database " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
		"data":    data,
	})

}

func ResponseDelete(rowsAffected int64, id string, c *gin.Context) {
	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success deleted data",
	})
}
