package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/urfave/cli/v2"
)

var CmdWeb = cli.Command{
	Name:   "web",
	Usage:  "Start Qwerty web server",
	Action: runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Value:   "3000",
			Aliases: []string{"p"},
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runWeb(c *cli.Context) error {
	engine := html.New("./internal/app/views", ".gohtml")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":3000"))
	return nil
}
