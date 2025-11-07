package builder

import "fmt"

func BuildProject(dir string) error {
	config, err := generateBuilderConfigFromJSON(fmt.Sprintf("%s/wool.json", dir))

	if err != nil {
		return err
	}

	fmt.Printf("Views: %s\n", config.ViewsPath)
	fmt.Printf("Components: %s\n", config.ComponentsPath)
	return nil
}
