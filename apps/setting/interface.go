package setting

type Service interface {
	DescribeSetting() (*Setting, error)
	UpdateSetting(*Setting) error
}
