package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(err error, c *gin.Context) {
	strError := err.Error()
	switch {
	case strings.Contains(strError, "no rows"):
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		c.Abort()
	case strings.Contains(strError, "check constraint"):
		var strErrorDetail string
		errorDetail := strings.Split(strError, "_")
		if len(errorDetail) == 4 {
			strErrorDetail = errorDetail[1] + " " + errorDetail[2]
		} else {
			strErrorDetail = errorDetail[1]
		}

		tableName := strings.Split(strings.Split(strError, ` "`)[1], `" `)[0]
		c.JSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Field %s is required for creating %s", strErrorDetail, tableName),
		})
		c.Abort()
	case strings.Contains(strError, "unique"):
		errorDetail := strings.Split(strError, "_")
		tableName := strings.Split(strings.Split(strError, ` "`)[1], `_`)[0]
		uppercaseTableName := strings.Replace(tableName, string(tableName[0]), strings.ToUpper(string(tableName[0])), 1)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("%s with that %s is already exist", uppercaseTableName, errorDetail[1]),
		})
		c.Abort()
	case strings.Contains(strError, "unregister"):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unregistered username or email",
		})
		c.Abort()
	case strings.Contains(strError, "incorrect"):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect username, email, or password",
		})
		c.Abort()
	}
}
