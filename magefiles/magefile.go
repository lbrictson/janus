//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	ResetTermColo = "\033[0m"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Builds the server binary to tmp/janus
func Build() error {
	mg.Deps(Generate)
	fmt.Println(Yellow + "Building server..." + ResetTermColo)
	cmd := exec.Command("go", "build", "-o", "tmp/janus", "./cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Generates the ent code
func Generate() error {
	fmt.Println(Yellow + "Running code gen..." + ResetTermColo)
	cmd := exec.Command("go", "generate", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Runs the tests
func Test() error {
	fmt.Println(Yellow + "Running tests..." + ResetTermColo)
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Runs the server
func Run() error {
	mg.Deps(Generate)
	mg.Deps(DockerUp)
	cmd := exec.Command("go", "run", "./cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Runs the server with dev reloading using air
func Dev() error {
	mg.Deps(Generate)
	mg.Deps(DockerUp)
	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Runs docker compose depdendencies
func DockerUp() error {
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Clean builds and artifacts
func Clean() {
	fmt.Println(Yellow + "Cleaning tmp files..." + ResetTermColo)
	os.RemoveAll("tmp/janus")
	os.RemoveAll("janus.db")
	os.RemoveAll("janus.db-shm")
	os.RemoveAll("janus.db-wal")
	os.RemoveAll("tmp/sessions")
	fmt.Println(BrightGreen + "Tmp files removed")
}

// Reset will destroy all local data, be careful
func Reset() {
	fmt.Println(Yellow + "Resetting database and containers, all data will be lost!" + ResetTermColo)
	cmd := exec.Command("docker", "compose", "down", "-v", "--rmi", "local")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(Red + "Failed to reset docker containers" + ResetTermColo)
		return
	}
	fmt.Println(BrightGreen + "Docker containers removed" + ResetTermColo)
	mg.Deps(Clean)
}

// AddModel adds a new model to the ent schema (ORM) and regenerates the code
func AddModel() {
	// Get argument
	if len(os.Args) < 2 {
		fmt.Println(Red + "No model specified" + ResetTermColo)
		fmt.Println(Red + "Usage: mage addModel <model>" + ResetTermColo)
		fmt.Println(Red + "Example: mage addModel User (Note that model is specified in pascal case" + ResetTermColo)
		return
	}
	model := os.Args[1]
	fmt.Println(Yellow + "Adding model: " + model + ResetTermColo)
	err := sh.RunV("go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", model)
	if err != nil {
		fmt.Println(Red + "Failed to add model" + ResetTermColo)
		return
	}
	fmt.Println(BrightGreen + "Model added, generating new code" + ResetTermColo)
	mg.Deps(Generate)
}
