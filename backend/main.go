package main

import (
	"net/http"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Allow frontend to access backend
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	router.POST("/calculate", func(c *gin.Context) {
		expr := c.PostForm("expression")

		evalExpr, err := govaluate.NewEvaluableExpression(expr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "Error"})
			return
		}

		result, err := evalExpr.Evaluate(nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": strconv.FormatFloat(result.(float64), 'f', -1, 64)})
	})

	router.Run(":8080")
}
