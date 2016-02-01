package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Printf("WELCOME!\n")
	quiz := initQuiz()
	printQuestions(quiz)
	fmt.Printf("Results: %d%\n", quiz.Result())
}

func initQuiz() Quiz {
	questions := generateFixtures()
	quiz := Quiz{questions}

	return quiz
}

func printQuestions(quiz Quiz) {
	var userAnswers = [2]string{}
	reader := bufio.NewReader(os.Stdin)
	for i,question := range quiz.Questions {
		fmt.Printf("%s\n", question.Content)
		for _,answer := range question.Answers {
			fmt.Printf("%s: %s\n", answer.Id, answer.Content)
		}

		userAnswer, _ := reader.ReadString('\n')
		userAnswers[i] = userAnswer
	}

	fmt.Printf("%v", userAnswers)
}

/**
 * Fixtures
 */
func generateFixtures() []Question {
	q1 := Question{
		Content: "Kamil's birth year?",
		Answers: []Answer{
			Answer{
				Id: "A",
				Content: "1992",
				Correct: false,
			},
			Answer{
				Id: "B",
				Content: "1988",
				Correct: true,
			},
			Answer{
				Id: "C",
				Content: "1995",
				Correct: false,
			},
			Answer{
				Id: "D",
				Content: "1986",
				Correct: false,
			},
		},
	}

	q2 := Question{
		Content: "Kamil's eyes color?",
		Answers: []Answer{
			Answer{
				Id: "A",
				Content: "Black",
				Correct: false,
			},
			Answer{
				Id: "B",
				Content: "Green",
				Correct: false,
			},
			Answer{
				Id: "C",
				Content: "Blue",
				Correct: true,
			},
			Answer{
				Id: "D",
				Content: "Grey",
				Correct: false,
			},
		},
	}

	return []Question{q1, q2}
}
