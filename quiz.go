package main

import(
	"bufio"
	"fmt"
	"encoding/csv"
	"os"
	"io"
	"strings"
	"flag"
)

func main(){

	timePtr := flag.Int("Time", 30, "an int")
	flag.Parse()
	fileName := askName()
	questions, answers, length := parseCSV(fileName)

	// parseCSV will return these values in the event that our inputs are invalid
	if (questions == nil && answers == nil && length == 0){
	} else {
		fmt.Printf("You have %d seconds!\n",*timePtr)
		runQuiz(questions, answers, length)
	}
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


// Parses through the CSV file then creates slices containing the questions and answers
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

			i++

			// If i exceeds the maximum number of questions, assumed to be 200, we can assume the file doesnt exist
			if (i > 200) {
				fmt.Println("Exceeded 200 questions and errors are occuring, your file likely doesnt exist!")
				return nil, nil, 0
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

// As per its name, display a question at a specified index
func displayQuestion(questions []string, i int ){

	// If the index is negative or the array doesn't exist, we can assume this has failed
	if i < 0 || questions == nil {
		return
	}

	fmt.Printf("%s : ", questions[i])
}

// Trims the newline character from the string obtained through the Reader
func trimNewline(str string) (string){
	// Check if the input is invalid
	if str == "" {
		fmt.Printf("Invalid string , returning nil\n")
		return ""
	}
	str = strings.TrimSuffix(str, "\n")
	return str
}

// Displays the current question then grabs the users answer and returns it
func displayQuiz(questions, answers []string, i int) (string){
	// Check if input is invalid
	if i < 0 || questions == nil || answers == nil{
		fmt.Printf("Invalid input\n")
		return ""
	}
	answerReader := bufio.NewReader(os.Stdin)
	displayQuestion(questions, i)
	currentAnswer, _ := answerReader.ReadString('\n')

	// Trims the newline character so we can properly compare it to the actual answer later
	currentAnswer = trimNewline(currentAnswer)
	return currentAnswer
}


// Actually runs the quiz and prints how many questions you got correct!
func runQuiz(questions, answers []string, length int){
	// Check if input is invalid
	if length < 0 || questions == nil || answers == nil{
		fmt.Printf("Invalid input\n")
		return
	}

	// Keeps track of how many questions were answered correctly
	correct := 0

	// Prints a fun starting message!
	fmt.Println("Starting Quiz!")
	fmt.Println("--------------")
	fmt.Println("Enter each of your answers after each question")

	for i:=0; i < length; i++{
		// Grab the user's answer to the current question
		currentAnswer := displayQuiz(questions, answers, i)

		if (answers[i] == currentAnswer){
			// if their answer is correct, make sure to increment our counter!
			fmt.Println("CORRECT!")
			correct++
		} else {
			// if it isn't, they got it wrong. Woeful!
			fmt.Println("WRONG!")
			}
		}

		fmt.Printf("You answered %d questions correctly out of %d questions\n", correct, len(questions))
	}
