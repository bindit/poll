package main

import "fmt"

func main() {
	fmt.Printf("WELCOMEs!")
	quiz := initQuiz()
	printQuestions(quiz)
	fmt.Printf("Results: %d%", quiz.Result())
}

func initQuiz() Quiz {
	questions := generateFixtures()
	quiz := Quiz{questions}

	return quiz
}

func printQuestions(quiz Quiz) {
	for _,question := range quiz.Questions {
		fmt.Printf("%s", question.Content)
	}
}

/**
 * Fixtures
 */
func generateFixtures() []Question {
	q1 := Question{
		Content: "Kamil's birth year",
		Answers: []Answer{
			Answer{
				Content: "1992",
				Correct: false,
			},
			Answer{
				Content: "1988",
				Correct: true,
			},
			Answer{
				Content: "1995",
				Correct: false,
			},
			Answer{
				Content: "1986",
				Correct: false,
			},
		},
	}

	q2 := Question{
		Content: "Kamil's girlfriend name",
		Answers: []Answer{
			Answer{
				Content: "Marta",
				Correct: false,
			},
			Answer{
				Content: "Agnieszka",
				Correct: false,
			},
			Answer{
				Content: "Anna",
				Correct: true,
			},
			Answer{
				Content: "Zofia",
				Correct: false,
			},
		},
	}

	return []Question{q1, q2}
}
