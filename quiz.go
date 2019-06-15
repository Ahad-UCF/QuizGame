package main

import(
	"bufio"
	"fmt"
	//"encoding/csv"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the name of your quiz CSV, if using default quiz simply press enter: ")
	text, _ := reader.ReadString('\n')
	if text == "\n"{
		text = ""
		fmt.Println("Using default quiz file.")
	} else {
		fmt.Printf("Using %s", text)
	}

}
