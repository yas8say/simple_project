package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Darth(c *router.Context, second, third string) {
	if second == "register" && third == "" && c.Method == "GET" {
		handleRegister(c)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		router.HandleCreateUserAutoForm(c, "")
		return
	}

	c.NotFound = true
}

func handleRegister(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}