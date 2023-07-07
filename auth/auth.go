package auth

var (
	Platform string = ""
	Password string = ""
)

func NewLogin(platform, password string) {
	// url := "http://localhost:5280/api/v1/login"
	// utils.HttpPostJson(url string, jsonStr []byte)
	Platform = platform
	Password = password
}
