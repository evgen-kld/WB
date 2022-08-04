package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func cd(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return err
	}
	return nil
}

func pwd() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func echo(str string) {
	fmt.Println(str)
}

func kill(id int) error {
	process, err := os.FindProcess(id)
	if err != nil {
		return err
	}
	return process.Kill()
}

func ps() error {
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

func quit() {
	os.Exit(0)
}

func Cmd(args []string) error {
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("the dir path is not specified")
		}
		return cd(args[1])
	case "pwd":
		return pwd()
	case "echo":
		if len(args) < 2 {
			echo("")
		} else {
			echo(args[1])
		}
	case "kill":
		if len(args) < 2 {
			return errors.New("id is not specified")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		return kill(id)
	case "ps":
		return ps()
	case "\\quit":
		quit()
	default:
		command := exec.Command(args[0], args[1:]...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		return command.Run()
	}
	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s$ ", path)

		input, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		args := strings.Split(strings.TrimSuffix(input, "\n"), " ")
		if len(args) < 1 {
			log.Println("incorrect input. Usage: {path}$ <command> [OPTIONS]")
		}

		if err := Cmd(args); err != nil {
			log.Println(err)
		}
	}
}
