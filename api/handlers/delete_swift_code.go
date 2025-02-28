package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) DeleteSwiftCode(c *gin.Context) {
	swiftCode := c.Param("swift-code")
	if len(swiftCode) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "swift code should have an exact length of 11",
		})
		return
	}

	_, err := handler.Store.DeleteSwiftCode(c, swiftCode)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "this entry does not exist",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot delete this entry - internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
