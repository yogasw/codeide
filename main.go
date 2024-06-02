package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
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
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	// Create a new command with the VS Code binary and the provided arguments
	vscodeCmd := exec.Command(
		vscodeBinaryDir,
		"--args",
		fmt.Sprintf("--user-data-dir=%s/config/%s/user-data", exPath, lang),
		fmt.Sprintf("--extensions-dir=%s/config/%s/extensions", exPath, lang),
	)

	// Run the command
	err = vscodeCmd.Run()
	if err != nil {
		fmt.Println("Error opening VS Code:", err)
	}

}
