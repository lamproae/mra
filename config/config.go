package config

type Config struct {
	DeviceName     string
	IP             string
	Port           string
	UserName       string
	Password       string
	EnablePrompt   string
	LoginPrompt    string
	PasswordPrompt string
	Prompt         string
	ModeDB         map[string]string
}
