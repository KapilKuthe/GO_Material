package model

type Email struct {
	From    string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Message string
}

type EmailConfig struct {
	UserName string
	Password string
	Host     string
}
