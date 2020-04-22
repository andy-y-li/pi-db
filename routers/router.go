package routers

import (
	"github.com/gin-gonic/gin"
	. "pi-db/apis"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/cpu")
	{
		v1.GET("/T", GetInfos)
	}
	return r
}
