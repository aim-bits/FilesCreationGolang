package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	content := "This is me learning GOLANG"

	file, err := os.Create("./test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("lenth is:", length)
	defer file.Close()  //defer is used incase there's need to continue writing other codes 

	readFile("./test.txt")
}

//This function is to read any file created
func readFile(filname string)  {
	databyte, err := os.ReadFile(filname)
	
	checkNilErr(err)

	fmt.Println("text data is in:", string(databyte))
}


// This function acts as a modifier to avoid long lines of error checking repetition 
func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}