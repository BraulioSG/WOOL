package builder

import (
	"encoding/json"
	"os"
)

type BuilderConfig struct {
	ViewsPath      string `json:"viewsDir"`
	ComponentsPath string `json:"componentsDir"`
}

func generateBuilderConfigFromJSON(jsonPath string) (BuilderConfig, error) {
	file, err := os.ReadFile(jsonPath)

	if err != nil {
		return BuilderConfig{}, err
	}

	var cb BuilderConfig = BuilderConfig{}

	json.Unmarshal(file, &cb)

	return cb, nil
}
