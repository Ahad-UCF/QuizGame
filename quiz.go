package main

import(
	"bufio"
	"fmt"
	"encoding/csv"
	"os"
	"io"
	"strings"
)
// TODO: Compomentalize code instead of running everything in main
func main(){

	fileName := askName()
	questions, answers, length := parseCSV(fileName)
	runQuiz(questions, answers, length)
}


func askName() (string){
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

	return csvName
}

func parseCSV(csvName string) ([]string, []string, int){
	// Initializing a counter variable to use later to store data from the CSV
	var i int = 0

	// Initializing a slice to store each question of the CSV
	questionStorage := make([]string, 0)

	// Initializing a slice to store each answer of the CSV
	answerStorage := make([]string, 0)

	// Open a file based on user input and then create a reader for it
	quizFile, _ := os.Open(csvName)
	csvReader := csv.NewReader(bufio.NewReader(quizFile))

	// Parses through the CSV
	for {
		// Read the output of each csv line and store them into variables line and error
		line, error := csvReader.Read()

		if error == io.EOF{
			// If we've reached the end of file, we're done parsing!
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
			questionStorage = append(questionStorage, line[0])
			answerStorage = append(answerStorage, line[1])
			i++
		}

	}

	return questionStorage, answerStorage, i
}

//TODO: add failsafes for the below functions incase any are nil!

// As per its name, display a question
func displayQuestion(questions []string, i int ){
	fmt.Printf("%s : ", questions[i])
}

/* Checks if a question is value. Returns 1 if it is and 0 otherwise
func checkAnswer(answers []string, i, answer int) (int){
	if answers[i] == string(answer){
		return 1
	} else {
		return 0
	}
}*/

func trimNewline(str string) (string){
	str = strings.TrimSuffix(str, "\n")
	return str
}

func displayQuiz(questions, answers []string, i int) (string){
	answerReader := bufio.NewReader(os.Stdin)
	displayQuestion(questions, i)
	currentAnswer, _ := answerReader.ReadString('\n')
	currentAnswer = trimNewline(currentAnswer)
	return currentAnswer
}

func runQuiz(questions, answers []string, length int){
	correct := 0
	fmt.Println("Starting Quiz!")
	fmt.Println("--------------")
	fmt.Println("Enter each of your answers after each question")
	for i:=0; i < length; i++{
		currentAnswer := displayQuiz(questions, answers, i)
		if (answers[i] == currentAnswer){
			fmt.Println("CORRECT!")
			correct++
		} else {
			fmt.Println("WRONG!")
			}
		}

		fmt.Printf("You answered %d questions correctly out of %d questions\n", correct, len(questions))
	}
