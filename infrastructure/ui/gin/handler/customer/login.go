package customer

import (
	"errors"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/jwt"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/presenter"
	"github.com/santosdavid/go-api-v2/usecase/customer/login"
)

func loginHandler(r *gin.Engine, service *login.Service) {
	r.POST("login", func(c *gin.Context) {
		var input struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			if reflect.TypeOf(err).Kind() == reflect.Slice {
				err := err.(validator.ValidationErrors).Translate(trans)

				c.JSON(
					http.StatusUnprocessableEntity,
					presenter.ErrorPresenter{Msg: "Invalid Entity", Errors: err},
				)
				return
			}

			c.IndentedJSON(http.StatusUnprocessableEntity, presenter.ErrorPresenter{Msg: err.Error()})
			return
		}

		resp, err := service.Login(input.Email, input.Password)
		if err != nil {
			if errors.Is(err, customer.ErrorInvalidPassword) || errors.Is(err, customer.ErrorEntityNotFound) {
				c.IndentedJSON(http.StatusUnauthorized, presenter.ErrorPresenter{Msg: "email or password invalid"})
				return
			}

			c.IndentedJSON(http.StatusInternalServerError, presenter.ErrorPresenter{Msg: "internal server error"})
			return
		}

		redis := storage.NewRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"))
		secrets := jwt.Secrets{[]byte(os.Getenv("JWT_SECRET")), []byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))}
		credentials, err := jwt.Generate(redis, secrets, &resp)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, presenter.ErrorPresenter{Msg: err.Error()})
		}

		c.IndentedJSON(http.StatusOK, credentials)
	})
}
