package customer

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/jwt"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"
)

func me(r *gin.RouterGroup) {
	r.GET("me", func(c *gin.Context) {
		c.JSON(http.StatusOK, "you")
	})
}

func checkLogin(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	redis := storage.NewRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"))
	secrets := jwt.Secrets{[]byte(os.Getenv("JWT_SECRET")), []byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))}

	_, err := jwt.ParseToken(redis, secrets, accessToken)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}
