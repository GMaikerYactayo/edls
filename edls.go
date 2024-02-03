package main

import (
	"github.com/fatih/color"
	"time"
)

const MacOS = "macOs"

const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleFileType struct {
	icon   string
	color  color.Attribute
	symbol string
}

var mapStyleFileType = map[int]styleFileType{
	fileRegular:    {icon: "📄"},
	fileDirectory:  {icon: "📁", color: color.BgHiBlue, symbol: "/"},
	fileExecutable: {icon: "🚀", color: color.BgGreen, symbol: "*"},
	fileCompress:   {icon: "📦", color: color.BgRed},
	fileImage:      {icon: "📸", color: color.BgMagenta},
	fileLink:       {icon: "🔗", color: color.BgCyan},
}

var (
	blue    = color.New(color.BgHiBlue).Add(color.Bold).SprintfFunc()
	green   = color.New(color.BgGreen).Add(color.Bold).SprintfFunc()
	red     = color.New(color.BgRed).Add(color.Bold).SprintfFunc()
	magenta = color.New(color.BgMagenta).Add(color.Bold).SprintfFunc()
	cyan    = color.New(color.BgCyan).Add(color.Bold).SprintfFunc()
	yellow  = color.New(color.BgYellow).SprintfFunc()
)
