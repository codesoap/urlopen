`urlopen` takes an URL as input and passes it to different programs,
depending on the type of the given URL. It is intended to be used with
tools like
[ddgr](https://github.com/jarun/ddgr)/[googler](https://github.com/jarun/googler),
[newsboat](https://github.com/newsboat/newsboat) or
[tuir](https://gitlab.com/ajak/tuir).

The type of URL is determined using regex patterns (see the `patterns`
constant in `main.go`). Feel free to create pull requests for new
patterns!

# Examples
There isn't much to show, but here are some examples with brief
explanations:
```console
$ # You could configure urlopen to play YouTube videos with mpv:
$ urlopen 'https://www.youtube.com/watch?v=gWsYaOM6pRg'

$ # You could configure urlopen to download and display pictures with
$ # wget and sxiv:
$ urlopen 'https://upload.wikimedia.org/wikipedia/commons/e/e1/Lada_Niva%2C_Sabir_%28P1090298%29.jpg'

$ # By default, Firefox is used as a fallback for opening unrecognized
$ # URLs:
$ urlopen 'https://en.wikipedia.org/wiki/Fukushima_Daiichi_nuclear_disaster'
```

# Installation
To install the binary to `~/go/bin/urlopen`, execute these steps (this
is using go modules, so it will only work with go version >=1.13, or
so):
```
git clone 'https://github.com/codesoap/urlopen.git'
cd urlopen
go install
```

# Configuration
To configure `urlopen`, you have to edit `config.go`, then rebuild and
install the program with `go install`. Examples and details are given in
the file.
