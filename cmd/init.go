package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Helper function to generate a new file with custom content.
func generateFile(fileName string, content string) {
	if content == "" {
		content = "#empty"
	}
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File exists\n")
	} else {
		d := []byte(content)
		check(ioutil.WriteFile(fileName, d, 0644))
	}

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long:  `A longer description of count command.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("initializing project...")

		name := viper.GetString("base.name")
		dir := viper.GetString("base.dir")

		os.Chdir(dir)
		var projectPath = path.Join(name)
		var appPath = path.Join(name, "app")
		var testPath = path.Join(name, "test")
		// fmt.Println(projectPath)

		err := os.MkdirAll(appPath, 0777)
		check(err)

		err = os.MkdirAll(testPath, 0777)
		check(err)

		// defer os.RemoveAll(appPath)

		touchProjectFiles := []string{".gitignore", ".dockerignore", "Dockerfile", "requirements.txt", "docker-compose.yml", "README.md", "LICENSE"}
		touchAppFiles := []string{"__init__.py", "main.py", "helpers.py"}
		touchTestFiles := []string{"__init__.py", "tests.py"}

		for _, i := range touchProjectFiles {
			generateFile(path.Join(projectPath, i), "")
		}

		for _, j := range touchAppFiles {
			generateFile(path.Join(appPath, j), "")
		}

		for _, j := range touchTestFiles {
			generateFile(path.Join(testPath, j), "")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
