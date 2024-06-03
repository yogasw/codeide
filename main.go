package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Simplifying VS Code for Open Specific Programming Languages",
	Run:   detectLanguage,
}

var vscodeBinaryDir string

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func detectLanguage(cmd *cobra.Command, args []string) {
	detectVSCodeBinary()

	if vscodeBinaryDir == "" {
		fmt.Println("error: VS Code binary not detected.")
		return

	}
	var lang string
	if len(args) > 0 {
		lang = args[0]
	}

	if lang == "create" {
		err := createProfile(args)
		if err != nil {
			fmt.Println("Error creating profile:", err)
			return
		} else {
			fmt.Println("Profile created successfully")
			return
		}
	}

	if lang == "" {
		fmt.Println("Detecting programming language...")
	}

	if lang == "golang" || lang == "go" {
		fmt.Println("Golang is detected")
	}

	openVSCode(cmd, lang)
}

func detectVSCodeBinary() {
	switch runtime.GOOS {
	case "darwin":
		vscodeBinaryDir = "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"
	}

}

func openVSCode(cmd *cobra.Command, lang string) {
	fmt.Println("Opening VS Code for " + lang)
	// Create a new command with the VS Code binary and the provided arguments
	var args []string

	if lang != "" {
		f := getConfigFolder(lang)
		if f == "" {
			fmt.Println("Please create profile for " + lang + " first.")
			fmt.Println("Run 'codeide create " + lang + "' to create profile.")
			return
		}
		args = append(args, "--args",
			fmt.Sprintf("--user-data-dir=%s/config/%s/user-data", f, lang),
			fmt.Sprintf("--extensions-dir=%s/config/%s/extensions", f, lang),
		)
	}

	vscodeCmd := exec.Command(
		vscodeBinaryDir,
		args...,
	)

	// Run the command
	err := vscodeCmd.Run()
	if err != nil {
		fmt.Println("Error opening VS Code:", err)
	}

}

func getHomeConfig() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	dir := filepath.Join(usr.HomeDir, ".codeide")

	if _, errC := os.Stat(dir); os.IsNotExist(errC) {
		errM := os.MkdirAll(dir, 0755)
		if errM != nil {
			panic(errM)
		}
	}
	configDir := filepath.Join(dir, "config")
	if _, errC := os.Stat(configDir); os.IsNotExist(errC) {
		errM := os.MkdirAll(configDir, 0755)
		if errM != nil {
			panic(errM)
		}
	}
	return configDir
}
func getConfigFolder(lang string) string {
	configDir := getHomeConfig()
	lDir := filepath.Join(configDir, lang)
	if _, errC := os.Stat(lDir); os.IsNotExist(errC) {
		return ""
	}

	return lDir
}

func createProfile(lang []string) error {
	configDir := getHomeConfig()
	if len(lang) < 2 {
		return fmt.Errorf("please provide a language to create profile, e.g. 'codeide create golang'")
	}

	lDir := filepath.Join(configDir, lang[1])
	if _, errC := os.Stat(lDir); os.IsNotExist(errC) {
		errM := os.MkdirAll(lDir, 0755)
		if errM != nil {
			return errM
		}
	} else {
		return fmt.Errorf("profile already exists")
	}

	return nil
}
