package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"url_base": "https://ekopras.engineer/api/v1",
		"url_auth": "https://ekopras.engineer/api/auth",
		"message":  "RESTful API : Golang (gin, gorm) x Mysql with JWT Auth",
		"alive":    true,
	})
}

func ResponseIndex(err error, data interface{}, c *gin.Context) {

	if err != nil {
		ResponseWithData(c, false, http.StatusInternalServerError, "Failed to retrieve data from the database "+err.Error(), []any{})
		return
	}

	ResponseWithData(c, true, http.StatusOK, "success", data)

}

func ResponseBindJson(err error, c *gin.Context) {
	if err != nil {
		Response(c, false, http.StatusBadRequest, "Failed to bind JSON "+err.Error())
		return
	}

}

func ResponseValidate(err error, c *gin.Context) {
	Response(c, false, http.StatusBadRequest, err.Error())
}

func ResponseCreate(err error, data interface{}, c *gin.Context) {

	if err != nil {
		Response(c, false, http.StatusInternalServerError, "Failed to create "+err.Error())
		return
	}

	Response(c, true, http.StatusOK, "create successfully")
}

func ResponseUpdate(rowsAffected int64, id string, data interface{}, c *gin.Context) {

	if rowsAffected == 0 {
		Response(c, false, http.StatusNotFound, "Record not found!")
		return
	}

	Response(c, true, http.StatusOK, "update successfully")
}

func ResponseShow(err error, data interface{}, c *gin.Context) {
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseWithData(c, false, http.StatusNotFound, "Record not found!", []any{})
			return
		default:
			Response(c, false, http.StatusInternalServerError, "Failed to retrieve data from the database "+err.Error())
			return
		}
	}

	ResponseWithData(c, true, http.StatusOK, "success", data)
}

func ResponseDelete(rowsAffected int64, id string, c *gin.Context) {
	if rowsAffected == 0 {
		Response(c, false, http.StatusBadRequest, "Record not found!")
		return
	}
	Response(c, true, http.StatusOK, "deleted successfully")
}

func Response(c *gin.Context, status bool, code int, message string) {
	c.JSON(code, gin.H{
		"status":  status,
		"code":    code,
		"message": message,
	})
	c.Abort()
}

func ResponseWithData(c *gin.Context, status bool, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"status":  status,
		"code":    code,
		"message": message,
		"data":    data,
	})
	c.Abort()
}
