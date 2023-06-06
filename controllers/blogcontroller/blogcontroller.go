package blogcontroller

import (
	"go-learn-restapi-mysql/config"
	"go-learn-restapi-mysql/controllers/base"
	"go-learn-restapi-mysql/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Index(c *gin.Context) {
	var blogs []models.Blog
	err := config.DB.Find(&blogs).Error
	if len(blogs) == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"status":  false,
			"message": "No data found",
			"data":    blogs,
		})
		return
	}
	// parse error and data to mastercontroller
	base.ResponseIndex(err, blogs, c)
}

func Create(c *gin.Context) {
	var blog models.Blog

	err := c.ShouldBindJSON(&blog)
	// parse error to mastercontroller
	base.ResponseBindJson(err, c)

	err = validate.Struct(blog)
	base.ResponseValidate(err, c)

	err = config.DB.Create(&blog).Error
	// parse error and data to mastercontroller
	base.ResponseCreate(err, blog, c)
}

func Show(c *gin.Context) {
	var blog models.Blog
	id := c.Param("id")

	err := config.DB.First(&blog, id).Error
	// parse error and data to mastercontroller
	base.ResponseShow(err, blog, c)

}

func Update(c *gin.Context) {
	var blog models.Blog
	id := c.Param("id")

	err := c.ShouldBindJSON(&blog)
	// parse error to mastercontroller
	base.ResponseBindJson(err, c)

	rowsAffected := config.DB.Model(&blog).Where("id = ?", id).Updates(&blog).RowsAffected
	// parse error and data to mastercontroller
	base.ResponseUpdate(rowsAffected, id, blog, c)
}

func Delete(c *gin.Context) {
	var blog models.Blog
	id := c.Param("id")

	rowsAffected := config.DB.Where("id = ?", id).Delete(&blog).RowsAffected
	// parse error and data to mastercontroller
	base.ResponseDelete(rowsAffected, id, c)
}
