package main

import (
	"fmt"
	"os"
)

const Version = "1.0.0"

// ANSI Colors
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Cyan   = "\033[36m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func showVersion() {
	fmt.Printf("\n%s⚡ xlr8%s %sv%s%s\n\n",
		Bold+Cyan,
		Reset,
		Green,
		Version,
		Reset,
	)
}

func showHelp() {
	fmt.Println(Bold + Cyan + `
╔══════════════════════════════════╗
║             ⚡ xlr8             ║
║      Terminal File Manager      ║
╚══════════════════════════════════╝
` + Reset)

	fmt.Println(Bold + "Usage" + Reset)
	fmt.Println("  xlr8")
	fmt.Println("  xlr8 --version")
	fmt.Println("  xlr8 --help")

	fmt.Println()

	fmt.Println(Bold + "Description" + Reset)
	fmt.Println("  Navigate to a target directory and run:")
	fmt.Println("  " + Green + "xlr8_go" + Reset)
	fmt.Println("  When you quit, the terminal updates to the selected path.")

	fmt.Println()

	fmt.Println(Bold + "Options" + Reset)
	fmt.Println("  " + Yellow + "-h, --help" + Reset + "      Show help")
	fmt.Println("  " + Yellow + "-v, --version" + Reset + "   Show version")

	fmt.Println()

	fmt.Println(Blue + "Written and directed by CS7player 🚀" + Reset)
	fmt.Println()
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			showVersion()
			return

		case "--help", "-h":
			showHelp()
			return
		}
	}

	StartScreen()
}
