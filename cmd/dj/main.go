package main

import (
	"fmt"
	"os"
	"time"

	"github.com/midry3/diajog/internal/diary"
)

var (
	args []string = os.Args[1:]
)

func getArg() string {
	v := args[0]
	args = args[1:]
	return v
}

func showHelp() {
	fmt.Println(`[Usage]
$ dj <option> "Content"

[Options]
-h, --help		Show help.
-v, --view		View your diary.`)
	os.Exit(0)
}

func view() {
	for day, info := range diary.GetDiary() {
		fmt.Printf("%s: %s\n", day, info.Content)
	}
	os.Exit(0)
}

func argparse() string {
	content := ""
	for 0 < len(args) {
		arg := getArg()
		switch arg {
		case "-h", "--help":
			showHelp()
		case "-v", "--view":
			view()
		default:
			content += arg + " "
		}
	}
	if content == "" {
		showHelp()
	}
	return content
}

func main() {
	content := argparse()
	diary.Record(content)
	fmt.Printf("Recorded '%s'\n", time.Now().Format("2006-01-02"))
}
