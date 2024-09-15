package main

import (
	"fmt"
	"os"
)

func main() {

	// - get file details
	// file, err := os.Open("23_files/example.txt")

	// if err != nil {
	// 	/// log the error
	// 	panic(err)
	// }

	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("File", file)
	// fmt.Println("File Info", fileInfo.Name())
	// fmt.Println("File or False", fileInfo.IsDir())
	// fmt.Println("File Size", fileInfo.Size())
	// fmt.Println("File Permission", fileInfo.Mode())
	// fmt.Println("File Modified At", fileInfo.ModTime())

	// - read file
	// f, err := os.Open("23_files/example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// file, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// buf := make([]byte, file.Size())

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("Data", d, string(buf[i]))
	// }

	// # Only for use small size file
	// data, err := os.ReadFile("23_files/example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// read folders
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// create file
	// f, err := os.Create("23_files/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("Faizan Shaikh")

	// bytes := []byte("Heloo Faizan Shaikh 2")

	// f.Write(bytes)

	// - read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("23_files/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// writerFile, err := os.Create("23_files/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer writerFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(writerFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(err)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("written to new file successfully")

	// Delete File

	err := os.Remove("23_files/example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")

}
