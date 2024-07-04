package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/darth/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "/aademo/start")
	}
	if Global.Start == "start.html" {
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "aademo", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "darth", nil, afterRegister)
	}
}