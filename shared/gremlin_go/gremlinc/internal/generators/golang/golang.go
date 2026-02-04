package golang

import (
	"os"
	"path/filepath"

	"github.com/norma-core/norma-core/shared/gremlin_go/gremlinc/internal/generators/golang/mapper"
	"github.com/norma-core/norma-core/shared/gremlin_go/gremlinc/internal/types"
)

func Generate(root string, modulePath string, targets []*types.ProtoFile) []error {
	mapped, errors := mapper.MapProtoFiles(root, modulePath, targets)
	if len(errors) > 0 {
		return errors
	}

	for _, target := range mapped {
		content := target.GenerateCode()
		_ = os.MkdirAll(filepath.Dir(target.FullOutputPath), 0755)
		if err := os.WriteFile(target.FullOutputPath, []byte(content), 0644); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
