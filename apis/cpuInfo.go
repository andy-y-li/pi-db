package apis

import (
	"github.com/gin-gonic/gin"
	. "pi-db/models"
)

// curl -i http://localhost:8080/api/cpu/T
func GetInfos(c *gin.Context) {
	infos := ListInfo()
	c.JSON(200, infos)
}
