package main

type Quiz struct {
	Questions []Question
	Answers []Answer
}

/**
 * Calculate Quiz results
 */
func (q Quiz) Result() int {
	return 0;
}
