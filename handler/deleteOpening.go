package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE opening",
	})
}
