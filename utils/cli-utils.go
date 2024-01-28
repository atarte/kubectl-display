package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// ClearScreen
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ExitLine
func ExitLine() {
	fmt.Println("")
	fmt.Println("Press [enter] to exit.")
	fmt.Println("")
}
