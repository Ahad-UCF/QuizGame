package main

import(
	"bufio"
	"fmt"
	"encoding/csv"
	"os"
	"io"
)

func main(){
	var i int = 0
	cmdReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the name of your quiz csv, if using default quiz simply press enter: ")
	text, _ := cmdReader.ReadString('\n')
	if text == "\n"{
		text = "Quiz.csv"
		fmt.Println("Using default quiz file.")
	} else {
		fmt.Printf("Using %s", text)
	}
	quizFile, _ := os.Open(text)
	csvReader := csv.NewReader(bufio.NewReader(quizFile))
	for {
		line, error := csvReader.Read()
		if error == io.EOF{
			fmt.Printf("Finished!\n")
			break
		} else if error != nil{
			fmt.Printf("%s", error)
		} else {
			i++
			fmt.Printf("%s", line)
		}
	}
}
