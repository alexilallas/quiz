package usecase

import (
	"github.com/alexilallas/quiz/internal/core/entity"
	"github.com/alexilallas/quiz/internal/core/port"
	"sync"
)

type Quiz struct {
}

const optionA = "A"
const optionB = "B"
const optionC = "C"

var questions entity.Questions
var answersAVG []float32
var m sync.Mutex

func init() {
	questions = []entity.Question{
		{
			Description: "How much is 1+1 ?",
			Options: map[string]*entity.Option{
				optionA: {
					Description: "2",
				},
				optionB: {
					Description: "10",
				},
				optionC: {
					Description: "It depends on the numeric base used in the calc.",
					IsCorrect:   true,
				},
			},
		},
		{
			Description: "What is the capital of Malta ?",
			Options: map[string]*entity.Option{
				optionA: {
					Description: "Silena",
				},
				optionB: {
					Description: "Valletta",
					IsCorrect:   true,
				},
				optionC: {
					Description: "Mdina",
				},
			},
		},
		{
			Description: "Do you like Golang ?",
			Options: map[string]*entity.Option{
				optionA: {
					Description: "Yes",
					IsCorrect:   true,
				},
				optionB: {
					Description: "No",
				},
				optionC: {
					Description: "I have tried once",
				},
			},
		},
	}
}

func (q Quiz) ListQuestions() entity.Questions {
	return questions
}

func (q Quiz) RegisterAnswers(answers entity.Answer) (r entity.Response) {
	m.Lock()
	defer m.Unlock()
	r.TotalQuestions = float32(len(questions))
	for key, question := range questions {
		if opt := question.Options[answers[key]]; opt != nil && opt.IsCorrect {
			r.Correct++
		}
	}

	avg := r.Correct / r.TotalQuestions
	var betterThan float32
	for _, ans := range answersAVG {
		if avg > ans {
			betterThan++
		}
	}

	if l := len(answersAVG); l > 0 {
		r.Percentage = betterThan / float32(l) * 100
	}

	answersAVG = append(answersAVG, avg)
	return
}

func ProvideQuizUseCase() port.QuizUseCase {
	return Quiz{}
}
