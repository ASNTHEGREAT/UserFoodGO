package middleware

import(
	"github.com/gin-gonic/gin"
	"github.com/ASNTHEGREAT/entity"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		entity.users.username : entity.users.password,
	})
}