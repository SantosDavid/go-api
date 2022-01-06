package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/usecase/customer/create"
	"github.com/santosdavid/go-api-v2/usecase/customer/login"
)

type Handler struct {
	repo customer.Repository
}

func newHandler(repo customer.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

var trans ut.Translator

func (h *Handler) Make(r *gin.Engine) {
	createService := create.New(h.repo)
	loginService := login.New(h.repo)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)
	}

	createHandler(r, createService)
	loginHandler(r, loginService)

	authorized := r.Group("")
	authorized.Use(checkLogin)
	me(authorized)
}
