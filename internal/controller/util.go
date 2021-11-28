package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

func HandleError(handler func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			switch val := err.(type) {
			case *errors.StatusError:
				c.JSON(val.Status, gin.H{"message": val.Err})
				break
			default:
				c.JSON(500, gin.H{"message": "internal server error: " + val.Error()})
			}
		}
	}
}

func GetInt(value string) int {
	if value == "" {
		return 0
	}
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int(val)
}
