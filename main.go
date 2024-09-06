package main

import (
	// "net/http"

	"math/rand"

	"github.com/corazawaf/coraza/v3"
	txhttp "github.com/corazawaf/coraza/v3/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()
	corazaCfg := coraza.NewWAFConfig().
		// WithDirectivesFromFile("./default.conf")
		WithDirectivesFromFile("./coraza.conf").
		WithDirectivesFromFile("./coreruleset/crs-setup.conf").
		WithDirectivesFromFile("./coreruleset/rules/*.conf")
		// WithDirectives(`SecRule ARGS "@rx <script>" "id:941100,phase:2,block,capture,t:none,msg:'XSS Attack Detected'"`)
	waf, err := coraza.NewWAF(corazaCfg)
	if err != nil {
		panic(err)
	}

	app.Get("/", adaptor.HTTPHandler(txhttp.WrapHandler(waf, adaptor.FiberHandlerFunc(phase1Trigger))))

	app.Listen(":8000")
}

func phase1Trigger(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"someResponsse": rand.New(rand.NewSource(2919)).Int(),
	})

}
