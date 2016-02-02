package main

type Quiz struct {
	Questions []Question
	Answers []Answer
}

/**
 * Calculate Quiz results
 */
func (q Quiz) Result() int {
	totalAnswers := float64(len(q.Answers))
	correctAnswers := float64(0)

	for _, answer := range q.Answers {
		if (answer.Correct) {
			correctAnswers++
		}
	}

	return int(correctAnswers / totalAnswers * 100);
}
