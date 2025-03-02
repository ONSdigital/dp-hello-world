package mapper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-hello-world-controller/config"
	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

// TODO: remove example test case
func TestUnitMapper(t *testing.T) {
	ctx := context.Background()

	Convey("test mapper adds emphasis to hello world string when set in config", t, func() {
		cfg := &config.Config{
			BindAddr:                   "1234",
			SiteDomain:                 "localhost",
			GracefulShutdownTimeout:    0,
			HealthCheckInterval:        0,
			HealthCheckCriticalTimeout: 0,
			PatternLibraryAssetsPath:   "http://localhost:9002/dist/assets",
			Debug:                      true,
			HelloWorldEmphasise:        true,
		}

		req := httptest.NewRequest("GET", "http://localhost/hello-world", http.NoBody)

		basePageModel := model.NewPage("", "localhost")

		hm := HelloModel{
			Greeting: "Hello",
			Who:      "World",
		}

		hw := CreateHelloWorldPage(ctx, req, cfg, basePageModel, hm, "en")
		So(hw.HelloWho, ShouldEqual, "Hello World!")
	})
}
