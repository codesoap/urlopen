package main

// commands contains a list of commands to execute, when the
// corresponding pattern matched the given URL. The first element in a
// command may be an environment variable. All possible keys are defined
// as constants in main.go. It is not necessary to provide a command for
// every key; when no match is found, the fallback command will be used.
//
// Examples:
//     youtubeVideo: []string{"mpv", "--fs", "--ytdl-format=bestvideo[height<=?1080]+bestaudio/best"},
//     png:          []string{"urlopen-png"},
//     fallback:     []string{"$BROWSER"},
var commands = map[match][]string{
	fallback: []string{"firefox"},
}
