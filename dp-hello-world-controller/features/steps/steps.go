package steps

import (
	"github.com/cucumber/godog"
)

func (c *Component) RegisterSteps(ctx *godog.ScenarioContext) {
	c.uiFeature.RegisterSteps(ctx)

	ctx.Step(`^the Hello World! element should be visible$`, c.theHelloWorldElementShouldBeVisible)
}

func (c *Component) theHelloWorldElementShouldBeVisible() error {
	// TODO: to build custom steps use uiFeature methods or alternatively use uiFeature.Chrome to directly drive the browser...
	return c.uiFeature.ElementShouldBeVisible("p.hello-world")
}
