package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Darth(c *router.Context, second, third string) {
	if second == "register" && third == "" && c.Method == "GET" {
		handleRegister(c)
		return
	}
	if second == "login" && third == "" && c.Method == "GET" {
		handleLogin(c)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		router.HandleCreateUserAutoForm(c, "")
		return
	}
	if second == "login" && third == "" && c.Method == "POST" {
		router.HandleCreateSessionAutoForm(c)
		return
	}
	if second == "logout" && third == "" && c.Method == "DELETE" {
		router.DestroySession(c)
		return
	}
	if router.NotLoggedIn(c) {
		return
	}
	if second == "start" && third == "" && c.Method == "GET" {
		handleStart(c)
		return
	}
	if second == "new-doc" && third == "" && c.Method == "GET" {
		handleNewDoc(c)
		return
	}

	c.NotFound = true
}

func handleRegister(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}
func handleLogin(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("login.html", send, 200)
}
func handleStart(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("start.html", send, 200)
}
func handleNewDoc(c *router.Context) {
	send := map[string]interface{}{}
	c.SendContentInLayout("new_doc.html", send, 200)
}
