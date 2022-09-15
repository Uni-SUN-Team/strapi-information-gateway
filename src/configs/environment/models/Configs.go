package models

type Configs struct {
	App struct {
		ENV         string `mapstructure:"env"`
		Port        string `mapstructure:"port"`
		ContextPath string `mapstructure:"context_path"`
	} `mapstructure:"app"`
	Gin struct {
		Mode     string `mapstructure:"mode"`
		RootPath string `mapstructure:"root_path"`
		Version  string `mapstructure:"version"`
		Configs  struct {
			TrustedProxies string `mapstructure:"trusted_proxies"`
		} `mapstructure:"configs"`
	} `mapstructure:"gin"`
	Swag struct {
		Title       string `mapstructure:"title"`
		Description string `mapstructure:"description"`
		Version     string `mapstructure:"version"`
		Host        string `mapstructure:"host"`
		Schemes     string `mapstructure:"schemes"`
	} `mapstructure:"swag"`
	Endpoint struct {
		Strapi struct {
			Host  string `mapstructure:"host"`
			Token string `mapstructure:"access_token"`
		} `mapstructure:"strapi"`
	} `mapstructure:"endpoint"`
}
