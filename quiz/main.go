package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Printf("WELCOME!\n")
	quiz := initQuiz()
	runQuiz(quiz)
	fmt.Printf("Results: %d%\n", quiz.Result())
}

func initQuiz() *Quiz {
	questions := generateFixtures()
	quiz := &Quiz{
		Questions: questions,
	}

	return quiz
}

func runQuiz(quiz *Quiz) {
	for _,question := range quiz.Questions {
		quiz.Answers = append(quiz.Answers, runQuestion(question))
	}
}

func runQuestion(question Question) Answer {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (Type an answer letter)\n", question.Content)

	for _,answer := range question.Answers {
		fmt.Printf("%s: %s\n", answer.Id, answer.Content)
	}

	answerId, _ := reader.ReadString('\n')
	for !question.IsAnswerDefined(strings.TrimSpace(strings.ToUpper(answerId))) {
		fmt.Printf("Incorrect answer! Try again:\n")
		answerId, _ = reader.ReadString('\n')
	}

	userAnswer, _ := question.GetAnswerById(strings.TrimSpace(strings.ToUpper(answerId)))

	return userAnswer
}

/**
 * Fixtures
 */
func generateFixtures() []Question {
	q1 := Question{
		Content: "Number of Star Wars episodes?",
		Answers: []Answer{
			Answer{
				Id: "A",
				Content: "3",
				Correct: false,
			},
			Answer{
				Id: "B",
				Content: "7",
				Correct: true,
			},
			Answer{
				Id: "C",
				Content: "6",
				Correct: false,
			},
			Answer{
				Id: "D",
				Content: "9",
				Correct: false,
			},
		},
	}

	q2 := Question{
		Content: "Real name of Darth Vader?",
		Answers: []Answer{
			Answer{
				Id: "A",
				Content: "Luke",
				Correct: false,
			},
			Answer{
				Id: "B",
				Content: "Kylo",
				Correct: false,
			},
			Answer{
				Id: "C",
				Content: "Anakin",
				Correct: true,
			},
			Answer{
				Id: "D",
				Content: "Ben",
				Correct: false,
			},
		},
	}

	return []Question{q1, q2}
}
