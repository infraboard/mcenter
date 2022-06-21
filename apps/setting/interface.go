package setting

type Service interface {
	GetSetting() (*Setting, error)
	UpdateSetting(*Setting) error
}
