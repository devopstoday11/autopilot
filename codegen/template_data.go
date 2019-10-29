package codegen

type TemplateData struct {
	Project

	TypesImportPrefix string // e.g. "v1"
	TypesImportPath   string // e.g. "github.com/solo-io/autopilot/examples/promoter/pkg/apis/canaries/v1"
	KindLowerCamel    string // e.g. "canary"
}