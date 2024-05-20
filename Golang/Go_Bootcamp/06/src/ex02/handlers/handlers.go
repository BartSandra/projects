package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"src/ex02/auth"
	"src/ex02/db"
	"src/ex02/models"
	"src/ex02/session"
	"strconv"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Limit(3).Offset((page - 1) * 3).Find(&posts)
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

func GetPost(c *gin.Context) {
	var post models.Post
	db.DB.Where("id = ?", c.Param("id")).First(&post)
	htmlContent := blackfriday.MarkdownCommon([]byte(post.Content))

	post.Content = string(htmlContent)

	c.HTML(http.StatusOK, "post.html", gin.H{
		"post": post,
	})
}

func AdminPanel(c *gin.Context) {
	session, _ := session.Store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}
	c.HTML(http.StatusOK, "admin_panel.html", nil)
}

func AdminLogin(c *gin.Context) {
	if !auth.CheckAdminCredentials(c) {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/admin")
	}
}

func CreatePost(c *gin.Context) {
	session, _ := session.Store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	htmlContent := blackfriday.MarkdownCommon([]byte(content))

	htmlContent = bytes.Replace(htmlContent, []byte("<p>"), []byte(""), -1)
	htmlContent = bytes.Replace(htmlContent, []byte("</p>"), []byte(""), -1)

	post := models.Post{Title: title, Content: string(htmlContent)}
	db.DB.Create(&post)

	c.Redirect(http.StatusFound, "/")
}

func EditPost(c *gin.Context) {
	session, _ := session.Store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	var post models.Post
	db.DB.Where("id = ?", c.Param("id")).First(&post)

	c.HTML(http.StatusOK, "admin_edit.html", gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	session, _ := session.Store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	htmlContent := blackfriday.MarkdownCommon([]byte(content))

	post := models.Post{Title: title, Content: string(htmlContent)}
	db.DB.Model(&models.Post{}).Where("id = ?", c.Param("id")).Updates(post)

	c.Redirect(http.StatusFound, "/")
}

func DeletePost(c *gin.Context) {
	session, _ := session.Store.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusUnauthorized, "admin_login.html", nil)
		return
	}

	db.DB.Where("id = ?", c.Param("id")).Delete(&models.Post{})

	c.Redirect(http.StatusFound, "/admin")
}

func SafeHTML(text string) template.HTML {
	return template.HTML(text)
}
