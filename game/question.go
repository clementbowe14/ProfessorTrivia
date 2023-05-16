package game

import (
	"errors"
	"strings"
)

const (
	MAX_QUESTION_LENGTH = 400
	MAX_ANSWER_LENGTH = 200
	MAX_HINT_LENGTH = 30
)
var (
	ErrorQuestionTooShort = "The question must contain at least 1 character"
	ErrorQuestionTooLong = "The question can not be longer than 400 characters"
	ErrorAnswerTooShort = "The answer must contain at least 1 character"
	ErrorAnswerTooLong = "The answer can not be longer than 200 characters"
	ErrorNoHintEntered = "Your hint must contain at least 1 character"
	ErrorHintMatchAnswer = "The hint can not be the same as the answer to the question"
	ErrorHintTooLong = "The hint can not exceed 30 characters in length"
)
type Question struct {
	points int
	topic string
	hint string
	query string
	answer string
}

func (q Question) GetQuestion() (string) {
	return q.query
}

func (q Question) GetAnswer() (string) {
	return q.answer
}

func (q Question) GetHint() (string) {
	return q.hint
}


func (q *Question) SetAnswer(answer string) (error) {
	if len(answer) == 0 {
		return errors.New(ErrorAnswerTooShort)
	}

	if len(answer) > MAX_ANSWER_LENGTH {
		return errors.New(ErrorAnswerTooLong)
	}

	q.answer = answer

	return nil

}

func (q *Question) SetHint(hint string) (error) {
	if len(hint) == 0 {
		return errors.New(ErrorNoHintEntered)
	}

	if len(hint) == len(q.answer) && strings.EqualFold(hint, q.answer) {
		return errors.New(ErrorHintMatchAnswer)
	}

	if len(hint) > MAX_HINT_LENGTH {
		return errors.New(ErrorHintTooLong)
	}

	return nil
}

func (q *Question) SetQuestion(question string) (error) {
	if len(question) == 0 {
		return errors.New(ErrorQuestionTooShort)
	}

	if len(question) > MAX_QUESTION_LENGTH {
		return errors.New(ErrorQuestionTooShort)
	}

	q.query = question

	return nil
}

