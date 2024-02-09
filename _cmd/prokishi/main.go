package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"prokishi"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/xerrors"
)

const iniFileName = "prokishi.ini"

// go build -ldflags "-X main.version=?"
var version string

type IniFile struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Code     string `toml:"code"`
	EngineId string `toml:"engineId"`
	Level    string `toml:"logLevel"`
}

var iniFile IniFile

func init() {
}

func main() {
	flag.Parse()
	err := run()
	if err != nil {
		msg := fmt.Sprintf("%+v", err)
		slog.Error(msg)
		fmt.Fprintf(os.Stderr, msg+"\n")
		os.Exit(1)
	}
}

func run() error {

	args := flag.Args()
	if len(args) >= 1 {
		sub := args[0]
		if sub == "version" {
			fmt.Println("prokishi version:%s", version)
			return nil
		}
	}

	err := loadIniFile()
	if err != nil {
		return xerrors.Errorf("loadIniFile() error: %w", err)
	}

	lv := parseLogLevel(iniFile.Level)
	defer prokishi.SetLogFile(lv, "prokishi", version == "").Close()

	err = prokishi.Run(iniFile.Host,
		iniFile.Port,
		prokishi.Code(iniFile.Code),
		prokishi.Engine(iniFile.EngineId),
		prokishi.Version(version))
	if err != nil {
		return xerrors.Errorf("prokishi.Run() error: %w", err)
	}
	return nil
}

func parseLogLevel(lv string) slog.Level {
	v := strings.ToLower(lv)
	switch v {
	case "dbg", "debug":
		return slog.LevelDebug
	case "info", "information":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "err", "error":
		return slog.LevelError
	default:
		return slog.LevelWarn
	}
	return slog.LevelInfo
}

func loadIniFile() error {

	//versionに値が入っているかで開発環境を判定
	dir, err := prokishi.GetRunDir(version == "")
	if err != nil {
		return fmt.Errorf("実行位置の取得に失敗しました")
	}

	p := filepath.Join(dir, iniFileName)
	if _, err := os.Stat(p); err != nil {
		return fmt.Errorf("%s が存在しません", p)
	}
	_, err = toml.DecodeFile(p, &iniFile)
	if err != nil {
		return fmt.Errorf("%s の解析に失敗しました", p)
	}

	return nil
}
