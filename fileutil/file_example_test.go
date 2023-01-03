package fileutil

import (
	"fmt"
	"os"
)

func ExampleIsExist() {

	result1 := IsExist("./")
	result2 := IsExist("./xxx.go")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleCreateFile() {
	fname := "./test.txt"

	result1 := IsExist(fname)

	CreateFile(fname)

	result2 := IsExist(fname)

	os.Remove(fname)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleCreateDir() {
	pwd, _ := os.Getwd()
	dirPath := pwd + "/test_xxx/"

	result1 := IsExist(dirPath)

	CreateDir(dirPath)

	result2 := IsExist(dirPath)

	os.Remove(dirPath)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleIsDir() {

	result1 := IsDir("./")
	result2 := IsDir("./xxx.go")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleRemoveFile() {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	copyFile := "./text_copy.txt"
	CopyFile(srcFile, copyFile)

	file, err := os.Open(copyFile)
	if err != nil {
		return
	}
	result1 := IsExist(copyFile)
	result2 := file.Name()

	os.Remove(srcFile)
	os.Remove(copyFile)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// ./text_copy.txt
}

func ExampleReadFileToString() {
	fname := "./test.txt"
	CreateFile(fname)

	f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	f.WriteString("hello world")

	content, _ := ReadFileToString(fname)

	os.Remove(fname)

	fmt.Println(content)

	// Output:
	// hello world
}

func ExampleClearFile() {
	fname := "./test.txt"
	CreateFile(fname)

	f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	f.WriteString("hello world")

	content1, _ := ReadFileToString(fname)

	ClearFile(fname)

	content2, _ := ReadFileToString(fname)

	os.Remove(fname)

	fmt.Println(content1)
	fmt.Println(content2)

	// Output:
	// hello world
	//
}

func ExampleReadFileByLine() {
	fname := "./test.txt"
	CreateFile(fname)

	f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	f.WriteString("hello\nworld")

	content, _ := ReadFileByLine(fname)

	os.Remove(fname)

	fmt.Println(content)

	// Output:
	// [hello world]
}

func ExampleListFileNames() {
	fileList, _ := ListFileNames("./")
	fmt.Println(fileList)

	// Output:
	// [file.go file_example_test.go file_test.go]
}

func ExampleZip() {
	srcFile := "./test.txt"
	CreateFile(srcFile)

	zipFile := "./test.zip"
	Zip(srcFile, zipFile)

	result := IsExist(zipFile)

	os.Remove(srcFile)
	os.Remove(zipFile)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleUnZip() {
	fname := "./test.txt"
	file, _ := os.Create(fname)
	file.WriteString("hello world")

	f, _ := os.Open(fname)
	defer f.Close()

	mimeType := MiMeType(f)
	fmt.Println(mimeType)

	os.Remove(fname)

	// Output:
	// application/octet-stream
}