package handler

import (
	"influencelab-backend/internal/model"
	"influencelab-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Analyze(c *gin.Context) {
	var req model.AnalyzeRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.Media == nil {
		mediaFile, err := c.FormFile("media")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "media file is required",
			})
			return
		}
		req.Media = mediaFile
	}

	result, err := service.AnalyzeContent(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
