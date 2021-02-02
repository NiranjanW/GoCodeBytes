package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := filepath.Join(".", "//temp")
	fmt.Println(dir)
	fp, _ := os.Create(filepath.Join(dir, "greeting.txt"))
	var sb strings.Builder
	greetingFile("Niran", fp)
	WriteGreetings("Nayan", &sb)
}

func greetingFile(msg string, f *os.File) error {

	greeting := Greeting(msg)
	_, err := f.Write([]byte(greeting))

	if err != nil {
		return err
	}
	return nil

}

func WriteGreeting(name string, sb *strings.Builder) error {
	greeting := Greeting(name)
	_, err := sb.Write([]byte(greeting))
	if err != nil {
		return err
	}
	return nil
}

func Greeting(name string) string {
	return fmt.Sprintf("Hello ,%v", name)
}

type Writer interface {
	Write([]byte) (int, error)
}

func WriteGreetings(name string, w Writer) error {
	greetings := Greeting(name)
	_, err := w.Write([]byte(greetings))
	if err != nil {
		return err
	}
	return nil
}

func ensureDir(dirName string) error {

	err := os.Mkdir(dirName, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
