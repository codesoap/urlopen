package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

type match int

const (
	// Websites:
	youtubeVideo = iota

	// Filetypes:
	gif
	jpg
	pdf
	png

	fallback
)

var patterns = map[match]string{
	youtubeVideo: `(?i)^(https://)?(www\.|m\.)?youtube\.[a-z]{2,4}/watch`,

	gif: `(?i)\.gif$`,
	jpg: `(?i)\.jpe?g$`,
	pdf: `(?i)\.pdf$`,
	png: `(?i)\.png$`,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Provide one URL as an argument.")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := openUrl(url); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the URL:", err.Error())
		os.Exit(2)
	}
}

func openUrl(url string) error {
	command := commands[fallback]
	var i match
	for i = 0; i < fallback; i++ {
		if pattern, ok := patterns[i]; ok {
			if !regexp.MustCompile(pattern).MatchString(url) {
				continue
			}
			var tmpCommand []string
			if tmpCommand, ok = commands[i]; ok {
				command = tmpCommand
				break
			}
		}
	}

	if len(commands) == 0 {
		return errors.New("command is empty")
	}
	cmd := exec.Command(os.ExpandEnv(command[0]))
	if len(commands) > 1 {
		cmd.Args = append(cmd.Args, command[1:]...)
	}
	cmd.Args = append(cmd.Args, url)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
