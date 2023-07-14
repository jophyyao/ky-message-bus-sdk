package kymessage

var (
	Platform string = ""
	Password string = ""
)

func NewLogin(platform, password string) {
	Platform = platform
	Password = password
}
