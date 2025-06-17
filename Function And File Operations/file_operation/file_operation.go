package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

const FilePerm = 0644
const DirPerm = 0755

func isDirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	// // File Info
	// info.Name()
	// info.Size()
	// info.ModTime()

	return info.IsDir()
}

func createDir(path string, nested bool) {
	var err error
	if nested {
		err = os.MkdirAll(path, DirPerm)
	} else {
		err = os.Mkdir(path, DirPerm)
	}
	if err != nil {
		panic(err)
	}
}

func getDir(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var dirs []string
	for _, filename := range files {
		if filename.IsDir() {
			dirs = append(dirs, filename.Name())
		}
		// // File Info
		// info, err := filename.Info()
		// if err != nil {
		// 	panic(err)
		// }
		// info.Name()
		// info.Size()
		// info.ModTime()
	}
	return dirs
}

func removeDir(path string, nested bool) {
	var err error
	if nested {
		err = os.RemoveAll(path)
	} else {
		err = os.Remove(path)
	}
	if err != nil {
		panic(err)
	}
}

func renameDir(oldPath string, newPath string) {
	if err := os.Rename(oldPath, newPath); err != nil {
		panic(err)
	}
}

func randomInt() int {
	return rand.Int()
}

func SaveDataToFileUsingTempFile(path string, data []byte) error {
	temp := fmt.Sprintf("%s.temp.%d", path, randomInt())
	file, err := os.OpenFile(temp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, FilePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		os.Remove(temp)
		return err
	}

	return os.Rename(temp, path)
}


func createFile(filename string) {
	file, err := os.Create(filename) // O_RDWR -> Read, Write
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func readFileIo(file *os.File) string {
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func readFileBufio(file *os.File) string {
	reader := bufio.NewReader(file)
	var content strings.Builder
	for {
		str, err := reader.ReadString('\n')
		content.WriteString(str)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
	return content.String()
}

func openAndRead(filename string) {
	file, err := os.Open(filename) // O_RDONLY -> Read
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content := readFileIo(file)
	content = readFileBufio(file)
	fmt.Println(content)
}

func appendToFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, FilePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
}

func overwriteFromBeginning(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, FilePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
}

func removeFile(filename string) {
	if err := os.Remove(filename); err != nil {
		panic(err)
	}
}

func main() {
}
