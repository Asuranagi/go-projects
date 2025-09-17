package api

type ApiService interface {
	Api()
}

type ApiServiceImpl struct{}

func NewApiService() ApiService {
	return &ApiServiceImpl{}
}

func (a *ApiServiceImpl) Api() {
}
