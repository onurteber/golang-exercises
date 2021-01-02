package main

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile  *os.File
	err      error
	fileInfo *os.FileInfo
)

func main() {

	// Create temp folder
	tempDirPath, err := ioutil.TempDir("", "tempFolder")
	if err != nil {
		log.Fatal(err)
	}
	println("Temp folder created. ", tempDirPath)

	// Create temp file
	tempFile, err := ioutil.TempFile(tempDirPath, "tempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	println("Temp file created. ", tempFile.Name())

	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}

	/*
		newFile, err = os.Create("demo.txt")
		if err != nil {
			log.Fatal(err)
		}

		/*fileInfo, err := os.Stat("demo.txt")
		if err != nil {
			log.Fatal(err)
		}

		println("File Name: ", fileInfo.Name())
		println("Size in bytes: ", fileInfo.Size())
		println("Permission: ", fileInfo.Mode())
		println("Last Modified Date: ", fileInfo.ModTime().String())
		println("Is Dictionary: ", fileInfo.IsDir())
		println("System Interface Type: ", fileInfo.Sys())
	*/

	/*
		// Move file
		originalPath := "demo.txt"
		newPath := "./moved/test.txt"
		errr := os.Rename(originalPath, newPath)
		if errr != nil {
			log.Fatal(errr)
			println(errr.Error())
		}
	*/

	/*
		fileInfo, err := os.Stat("demo.txt")
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatal("File does not exist!")
			}
		}
		println("File Info: ", fileInfo)
	*/
	/*
			// File open and close
		file, err := os.OpenFile("demo.txt", os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
		}
		// Write data
		file.Close()
	*/
	/*
		file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				log.Fatal(err)
			}
		}
		file.Close()
	*/
	/*
			// Copy File
		originalFile, err := os.Open("demo.txt")
		if err != nil {
			log.Fatal(err)
		}

		newFile, err := os.Create("./moved/demo_copy.txt")
		if err != nil {
			log.Fatal(err)
		}

		byteWritten, err := io.Copy(newFile, originalFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Copied %d bytes. ", byteWritten)

		err = newFile.Sync()
		if err != nil {
			log.Fatal(err)
		}
		originalFile.Close()
	*/
	/*
		// Write Bytes to a File

		file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		byteSlice := []byte("test\ntest2")
		byteWritten, err := file.Write(byteSlice)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Wrote %d bytes.\n", byteWritten)
		file.Close()
	*/

}
