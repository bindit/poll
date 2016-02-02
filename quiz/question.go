package main

import "errors"

type Question struct {
	Content string
	Answers []Answer
}

func (q Question) GetAnswerById(id string) (Answer, error) {
	for _,answer := range q.Answers {
		if answer.Id == id {
			return answer, nil
		}
	}

	return Answer{}, errors.New("Answer do not exist.")
}

func (q Question) IsAnswerDefined(id string) bool {
	_, err := q.GetAnswerById(id)
	return err == nil
}
