package mapper

import (
	"context"
	"github.com/ONSdigital/dp-hello-world-controller/config"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {
	ctx := context.Background()

	Convey("test CreateFilterOverview correctly maps item to filterOverview page model", t, func() {
		cfg := config.Config{
			BindAddr:                   "1234",
			GracefulShutdownTimeout:    0,
			HealthCheckInterval:        0,
			HealthCheckCriticalTimeout: 0,
			HelloWorldEmphasise:        true,
		}

		hm := HelloModel{
			Greeting: "Hello",
			Who:      "World",
		}

		hw := HelloWorld(ctx, hm, cfg)
		So(hw, ShouldEqual, "Hello World!")
	})
}
