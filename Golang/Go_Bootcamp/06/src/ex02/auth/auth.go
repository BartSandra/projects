package auth

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"src/ex02/session"
	"strings"
)

func CheckAdminCredentials(c *gin.Context) bool {
	session, err := session.Store.Get(c.Request, "session-name")
	if err != nil {
		log.Printf("Ошибка получения сессии: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return false
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	adminCredentials, err := ioutil.ReadFile("admin_credentials.txt")
	if err != nil {
		log.Printf("Ошибка чтения файла admin_credentials.txt: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return false
	}

	credentials := strings.Split(string(adminCredentials), "\n")
	if len(credentials) < 2 {
		log.Println("Файл admin_credentials.txt имеет неправильный формат.")
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return false
	}

	usernameFromFile := strings.TrimSpace(credentials[0])
	passwordFromFile := strings.TrimSpace(credentials[1])

	if username == usernameFromFile && password == passwordFromFile {
		session.Values["authenticated"] = true
		err = session.Save(c.Request, c.Writer)
		if err != nil {
			log.Printf("Ошибка сохранения сессии: %v", err)
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return false
		}
		log.Println("Аутентификация прошла успешно")
		return true
	} else {
		log.Printf("Неудачная попытка входа с username: %s", username)
		return false
	}
}
