package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() // create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows example, its tested
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

// Clear runs the command to clear the terminal
func Clear() {
	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
