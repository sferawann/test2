package middleware

import (
	"fmt"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test2/config"
	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/repository"
	"github.com/sferawann/test2/utils"
)

func DeserializeBorrower(borrowerRepository repository.BorrowerRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err := strconv.ParseInt(fmt.Sprint(sub), 10, 64)
		helper.ErrorPanic(err)
		result, err := borrowerRepository.FindById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		c.Set("currentUser", result.Username)
		c.Next()

	}
}
