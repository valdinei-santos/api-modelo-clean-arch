package controller

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string) error
}
