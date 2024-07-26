package presenter

// IView - Output
type IView interface {
	Show(stamp string, out *Output) error
	ShowError(stamp string, msgErro string) error
}
