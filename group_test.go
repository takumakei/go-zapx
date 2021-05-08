package zapx_test

import (
	"github.com/takumakei/go-zapx"
	"go.uber.org/zap"
)

func ExampleG() {
	log := zap.NewExample()
	log.Info(
		"example of using group",
		zapx.G(
			"groupA",
			zap.String("name", "alice"),
			zap.Int("star", 42),
		),
	)
	// Output:
	// {"level":"info","msg":"example of using group","groupA":{"name":"alice","star":42}}
}
