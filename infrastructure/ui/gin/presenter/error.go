package presenter

type ErrorPresenter struct {
	Msg    string      `json:"message"`
	Errors interface{} `json:"errors"`
}
