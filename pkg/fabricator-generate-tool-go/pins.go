// Code generated by fabricator-generate-plugin-go
//
// Modifications in code regions will be lost during regeneration!

package fabricatorgeneratetoolgo

// CODE_REGION(PINS)

var DefaultPins PinDependencies = PinDependencies{
	"go.opentelemetry.io/otel": {
		Name:    "go.opentelemetry.io/otel",
		Version: "v0.20.0",
	},
}

var DefaultReplacements ReplaceDependencies = ReplaceDependencies{
	// "github.com/onsi/ginkgo": {
	// 	Name: "github.com/onsi/ginkgo",
	// 	With: "github.com/magicmoose/ginkgo@v1.17.0",
	// },
}

var DefaultToolDependencies ToolDependencies = ToolDependencies{
	"github.com/onsi/ginkgo": {
		Name: "github.com/onsi/ginkgo/ginkgo",
	},
	"github.com/google/wire": {
		Name: "github.com/google/wire",
	},
	"github.com/google/wire/cmd/wire": {
		Name: "github.com/google/wire/cmd/wire",
	},
}

//endregion
