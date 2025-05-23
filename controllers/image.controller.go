package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/middlewares"
	"github.com/lubie-placki-be/services"
)

func UploadImage(c *gin.Context) {
	if !middlewares.IsAuthenticated {
		c.IndentedJSON(http.StatusUnauthorized, nil)
	}

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	fileId, err := services.UploadImage(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": fileId})
}

func DownloadImage(c *gin.Context) {
	id := c.Param("id")
	buf, err := services.DownloadImage(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	contentType := http.DetectContentType(buf.Bytes())

	c.Writer.Header().Add("Content-Type", contentType)
	c.Writer.Header().Add("Content-Length", strconv.Itoa(len(buf.Bytes())))
	c.Writer.Header().Set("Cache-Control", "max-age=36000, public")

	c.Writer.Write(buf.Bytes())
}
