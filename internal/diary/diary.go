package diary

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	APPNAME string = "diajog"
)

var (
	DIARY_PATH string = getDiaryPath()
)

type DiaryInfo struct {
	Content string
}

func showUnableGetPathErr() {
	log.Fatal("Faild to get the path. Please check your environment variables.")
}

func getDiaryPath() string {
	var base string

	switch runtime.GOOS {
	case "windows":
		base = os.Getenv("APPDATA")
		if base == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				showUnableGetPathErr()
			}
			base = filepath.Join(home, "Library", "Application Support")
		}

	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			showUnableGetPathErr()
		}
		base = filepath.Join(home, "Library", "Application Support")

	default:
		base = os.Getenv("XDG_DATA_HOME")
		if base == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				showUnableGetPathErr()
			}
			base = filepath.Join(home, ".local", "share")
		}
	}

	dir := filepath.Join(base, APPNAME)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Printf("Faild to create directory to '%s'\n", dir)
		os.Exit(1)
	}
	f := filepath.Join(dir, "diary.json")
	_, err := os.ReadFile(f)
	if err != nil {
		os.WriteFile(f, []byte("{}"), 0644)
	}
	return f
}

func readDiary() []byte {
	d, err := os.ReadFile(DIARY_PATH)
	if err != nil {
		fmt.Printf("Cannot read diary: '%s'\n", DIARY_PATH)
		os.Exit(1)
	}
	return d
}

func GetDiary() map[string]DiaryInfo {
	var diary map[string]DiaryInfo
	json.Unmarshal(readDiary(), &diary)
	return diary
}

func Record(txt string) {
	day := time.Now().Format("2006-01-02")
	info := DiaryInfo{
		Content: txt,
	}
	diary := GetDiary()
	diary[day] = info
	b, err := json.Marshal(diary)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile(DIARY_PATH, b, 0644); err != nil {
		log.Fatal(err)
	}
}
