package main

import (
	"aademo/app"
	"embed"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]

	if arg == "import" {
		// Implement import functionality if needed
	} else if arg == "render" {
		//router.RenderMarkup()
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)

		r.Paths["/"] = app.Welcome   // Adding this line
		r.Paths["darth"] = app.Darth // Adding this line
		r.Paths["markup"] = router.Markup
		//r.Paths["aademo_test"] = app.aademo_test
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
		// Provide help information if needed
	}
}
