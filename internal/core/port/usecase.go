package port

import "github.com/alexilallas/quiz/internal/core/entity"

type QuizUseCase interface {
	ListQuestions() entity.Questions
	RegisterAnswers(entity.Answer) entity.Response
}
