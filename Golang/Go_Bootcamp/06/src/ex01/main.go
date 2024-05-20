package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	gorm.Model
	Title   string
	Content string
}

var DB *gorm.DB
var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=db password=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	DB.AutoMigrate(&Post{})

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"safeHTML": safeHTML,
	})

	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", getPosts)
	r.GET("/post/:id", getPost)
	r.GET("/admin", adminPanel)
	r.POST("/admin", adminLogin)
	r.POST("/admin/post", createPost)
	r.GET("/admin/edit/:id", editPost)
	r.POST("/admin/edit/:id", updatePost)
	r.POST("/admin/delete/:id", deletePost)

	r.Static("/web/css/", "./web/css/")
	r.Static("/web/images", "./web/images")

	r.Run(":8888")
}

func safeHTML(text string) template.HTML {
	return template.HTML(text)
}

func deletePost(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	DB.Where("id = ?", c.Param("id")).Delete(&Post{})

	c.Redirect(http.StatusFound, "/admin")
}

func editPost(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	var post Post
	DB.Where("id = ?", c.Param("id")).First(&post)

	c.HTML(http.StatusOK, "admin_edit.html", gin.H{
		"post": post,
	})
}

func updatePost(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	htmlContent := blackfriday.MarkdownCommon([]byte(content))

	post := Post{Title: title, Content: string(htmlContent)}
	DB.Model(&Post{}).Where("id = ?", c.Param("id")).Updates(post)

	c.Redirect(http.StatusFound, "/")
}

func getPosts(c *gin.Context) {
	var posts []Post
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Limit(3).Offset((page - 1) * 3).Find(&posts)
	for i, post := range posts {
		if len(post.Content) > 100 {
			posts[i].Content = post.Content[:200] + "..."
		}
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"posts":    posts,
		"nextPage": page + 1,
	})
}

func getPost(c *gin.Context) {
	var post Post
	DB.Where("id = ?", c.Param("id")).First(&post)
	htmlContent := blackfriday.MarkdownCommon([]byte(post.Content))

	post.Content = string(htmlContent)

	c.HTML(http.StatusOK, "post.html", gin.H{
		"post": post,
	})
}

func adminPanel(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}
	c.HTML(http.StatusOK, "admin_panel.html", nil)
}

func adminLogin(c *gin.Context) {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		log.Printf("Ошибка получения сессии: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	//log.Printf("Полученные данные из формы: username=%s\n, password=%s\n", username, password)

	adminCredentials, err := ioutil.ReadFile("admin_credentials.txt")
	if err != nil {
		log.Printf("Ошибка чтения файла admin_credentials.txt: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	//log.Printf("Содержимое файла admin_credentials.txt: %s\n", string(adminCredentials))

	credentials := strings.Split(string(adminCredentials), "\n")
	if len(credentials) < 2 {
		log.Println("Файл admin_credentials.txt имеет неправильный формат.")
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	//log.Printf("Форма отправлена с username: '%s'\n и password: '%s'\n", username, password)
	//log.Printf("Учетные данные из файла: username: '%s'\n, password: '%s'\n", credentials[0], credentials[1])
	//log.Printf("Учетные данные для проверки: username=%s\n, password=%s\n", credentials[0], credentials[1])

	usernameFromFile := strings.TrimSpace(credentials[0])
	passwordFromFile := strings.TrimSpace(credentials[1])

	if username == usernameFromFile && password == passwordFromFile {
		session.Values["authenticated"] = true
		err = session.Save(c.Request, c.Writer)
		if err != nil {
			log.Printf("Ошибка сохранения сессии: %v", err)
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		log.Println("Аутентификация прошла успешно")
		c.Redirect(http.StatusFound, "/admin")
	} else {
		log.Printf("Неудачная попытка входа с username: %s", username)
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
	}
}

func createPost(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	htmlContent := blackfriday.MarkdownCommon([]byte(content))

	htmlContent = bytes.Replace(htmlContent, []byte("<p>"), []byte(""), -1)
	htmlContent = bytes.Replace(htmlContent, []byte("</p>"), []byte(""), -1)

	post := Post{Title: title, Content: string(htmlContent)}
	DB.Create(&post)

	c.Redirect(http.StatusFound, "/")
}
