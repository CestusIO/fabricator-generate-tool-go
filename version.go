// Code generated by fabricator-generate-plugin-go
//
// Modifications in code regions will be lost during regeneration!

package fabricatorgeneratetoolgo

// region CODE_REGION(version)
import (
	_ "embed"

	"code.cestus.io/libs/buildinfo"
)

//go:embed version.yml
var version string

func init() {
	buildinfo.GenerateVersionFromVersionYaml(GetVersionYaml(), "fabricator-generate-tool-go")
}

func GetVersionYaml() []byte {
	return []byte(version)
}

//endregion