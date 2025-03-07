package generate

import (
	"fmt"

	"github.com/JsonLee12138/jsonix/embed"
	"github.com/JsonLee12138/jsonix/pkg/core"
	"github.com/JsonLee12138/jsonix/pkg/utils"
)

func GenerateEntity(entityName, outPath string, override bool) error {
	return utils.TryCatchVoid(func() {
		tmpl := string(utils.Raise(embed.TemplatesPath.ReadFile("templates/entity.tmpl")))
		entityPath := fmt.Sprintf("%s/entity/%s.go", outPath, entityName)
		upperName := utils.UpperCamelCase(entityName)
		params := map[string]string{
			"Name": upperName,
		}
		utils.RaiseVoid(core.GenerateFileExistsHandler(entityPath, tmpl, params, override))
		fmt.Printf("✅ '%s' entity has been successfully generated!\n", entityName)
	}, utils.DefaultErrorHandler)
}
