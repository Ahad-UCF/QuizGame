# Why this repo exists
This repo serves as a way to hold my answer to the first question on the Gophercises website. I will continue to answer additional questions and upload my answers on this github page as I learn both git and golang.

## QuizGame
This golang document will read questions (with a limit of 200 questions) from a csv document and then allow the user to input their answers within the provided time limit. Following either the end of the time limit or the completion of all questions, the amount of correct answers by the user will be printed to the console. Note that the quiz will end as soon as the time limit is reached, even if the user is still typing an answer.

### How to run
In order to run this program from a terminal, enter the following command:
go build quiz.go &&./quiz -Time= <Your Desired Time Limit>
Where <Your Desired Time Limit> should be the amount of seconds you wish to set as your time limit. Note that this should be entered as an integer value rather than a string. In other words, the value 5 should be entered rather than the word five.
