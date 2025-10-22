package api

type Api struct {
	BaseApi
	UserApi
}

var ApiApp = new(Api)
