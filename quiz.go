package main

import(
	"bufio"
	"fmt"
	"encoding/csv"
	"os"
	"io"
)
// TODO: Compomentalize code instead of running everything in main
func main(){
	// Initializing a counter variable to use later to store data from the CSV
	var i int = 0

	// Initializing a slice to store each line of the CSV
	lineStorage := make([]string, 0)

	// This will read the command line to grab the name of the CSV file
	cmdReader := bufio.NewReader(os.Stdin)

	// Print a prompt and then sotre the name os the csv into a string
	fmt.Printf("Enter the name of your quiz csv, if using default quiz simply press enter: ")
	csvName, _ := cmdReader.ReadString('\n')

	if csvName == "\n"{
		// Use the default file if the user has chosen to
		csvName = "Quiz.csv"
		fmt.Println("Using default quiz file.")
	}  else {
		// set the csvName variable to the user input
		fmt.Printf("Using %s", csvName)
	}

	// Open a file based on user input and then create a reader for it
	quizFile, _ := os.Open(csvName)
	csvReader := csv.NewReader(bufio.NewReader(quizFile))

	// Parses through the CSV
	for {
		// Read the output of each csv line and store them into variables line and error
		line, error := csvReader.Read()
		
		if error == io.EOF{
			// If we've reached the end of file, we're done parsing!
			fmt.Printf("Finished!\n")
			break
		} else if error != nil{
			// If an error has been produced, print it out and continue incrementing i
			fmt.Println("%s", error)

			// If i exceeds the maximum number of questions, assumed to be 200, we can assume the file doesnt exist
			i++
			if (i > 200) {
				fmt.Println("Exceeded 200 questions and errors are occuring, your file likely doesnt exist!")
				break
			}
		} else {
			// Store the line into the slice and then increment the counter
			lineStorage = append(lineStorage, line[0])
			i++
		}
	}
}
