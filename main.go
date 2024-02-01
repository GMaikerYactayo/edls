package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func main() {
	// filter patter
	flagPattern := flag.String("p", "", "filter by pattern")
	//flagAll := flag.String("a", "", "all file including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of records")

	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []file{}
	for _, dir := range dirs {
		f, err := getFile(dir, false)
		if err != nil {
			panic(err)
		}

		isMatched, err := regexp.MatchString("(?i)"+*flagPattern, f.name)
		if err != nil {
			panic(err)
		}

		if !isMatched {
			continue
		}

		fs = append(fs, f)
	}
	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}
	printList(fs, *flagNumberRecords)

}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleFileType[file.fileType]
		fmt.Printf("%s %s %s %10d %s %s %s%s\n", file.mode, file.userName,
			file.groupName, file.size, file.modificationTime.Format(time.DateTime), style.icon, file.name, style.symbol)
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}
	f := file{
		name:             dir.Name(),
		isDir:            info.IsDir(),
		isHidden:         isHidden,
		userName:         "Maiker",
		groupName:        "Gonzales",
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}
	setFile(&f)
	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func isLink(f file) bool {
	return strings.HasPrefix(f.name, "L")
}

func isExec(f file) bool {
	if runtime.GOOS == MacOS {
		return strings.HasSuffix(f.name, exe)
	}
	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}
