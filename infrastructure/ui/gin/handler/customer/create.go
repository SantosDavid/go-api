package customer

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/presenter"
	"github.com/santosdavid/go-api-v2/usecase/customer/create"
)

func createHandler(r *gin.Engine, service *create.Service) {

	r.POST("customers", func(c *gin.Context) {
		var input struct {
			PayDay   int    `json:"payday" binding:"required,numeric"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=6"`
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

			c.JSON(
				http.StatusUnprocessableEntity,
				presenter.ErrorPresenter{Msg: "Invalid Entity", Errors: err},
			)
			return
		}

		req := create.Request{
			PayDay:   input.PayDay,
			Email:    input.Email,
			Password: input.Password,
		}

		customerID, err := service.Create(req)
		if err != nil {
			if errors.Is(err, customer.ErrorDuplicatedEmail) || errors.Is(err, customer.ErrorInvalidPayDay) {
				c.JSON(http.StatusUnprocessableEntity, presenter.ErrorPresenter{Msg: err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, presenter.ErrorPresenter{Msg: "internal server error"})
			return
		}

		c.JSON(http.StatusCreated, presenter.CustomerID{ID: customerID})
	})
}
