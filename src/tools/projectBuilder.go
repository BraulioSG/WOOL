package tools

import (
	"fmt"
	"os"
	"wool/builder"
)

func createProjectStructure(dir string) error {
	const perms os.FileMode = os.FileMode(0755) //file permissions: Owner, Group, Others can read, write and execute

	//create root folder
	var err error = os.MkdirAll(dir, perms)
	if err != nil {
		return fmt.Errorf("[ERROR] couldn't create new project %s", err.Error())
	}

	//create views folder
	err = os.MkdirAll(fmt.Sprintf("%s/views", dir), perms)
	if err != nil {
		return fmt.Errorf("[ERROR] couldn't create new project %s", err.Error())
	}

	//crate components folder
	err = os.MkdirAll(fmt.Sprintf("%s/components", dir), perms)
	if err != nil {
		return fmt.Errorf("[ERROR] couldn't create new project %s", err.Error())
	}

	return nil
}

func generateJsonFile(file *os.File) {

}

func CrateProject(dir string) (bool, error) {
	//const perms os.FileMode = os.FileMode(0755) //file permissions: Owner, Group, Others can read, write and execute

	var err error = createProjectStructure(dir)

	if err != nil {
		return false, err
	}

	//adding the files
	file, err := os.Create(fmt.Sprintf("%s/wool.json", dir))

	if err != nil {
		return false, fmt.Errorf("[ERROR] couldn't create config file")
	}

	generateJsonFile(file)

	return true, nil
}

func validateProjectStructure(dir string) error {
	if !isDir(dir) {
		return fmt.Errorf("[ERROR] %s directory not found", dir)
	}

	if !isFile(fmt.Sprintf("%s/wool.json", dir)) {
		return fmt.Errorf("[ERROR] %s/wool.json not found", dir)
	}

	return nil
}

func Build(dir string) (bool, error) {
	var err error = validateProjectStructure(dir)

	if err != nil {
		return false, err
	}

	err = builder.BuildProject(dir)

	if err != nil {
		return false, err
	}

	return true, nil
}
