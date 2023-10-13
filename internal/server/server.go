package server

import (
	"context"
	"github.com/alexilallas/quiz/internal/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"sync"
)

const optionA = "A"
const optionB = "B"
const optionC = "C"

type Server struct {
	grpc.UnimplementedQuizServer
	validator Validator
}

var questions grpc.Questions
var answersAVG []float32
var m sync.Mutex

func init() {
	questions = grpc.Questions{
		Questions: []*grpc.Question{
			{
				Description: "How much is 1+1 ?",
				Options: map[string]*grpc.Option{
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
				Options: map[string]*grpc.Option{
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
				Options: map[string]*grpc.Option{
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
		}}
}

func (Server) ListQuestions(context.Context, *empty.Empty) (*grpc.Questions, error) {
	return &questions, nil
}

func (s Server) RegisterAnswers(_ context.Context, a *grpc.Answer) (*grpc.QuizResponse, error) {
	var r = new(grpc.QuizResponse)
	if err := s.validator.Request(a); err != nil {
		return r, err
	}

	s.calculate(r, a.Answer)

	return r, nil
}

func (s Server) calculate(r *grpc.QuizResponse, answers []string) {
	m.Lock()
	defer m.Unlock()
	r.TotalQuestions = float32(len(questions.Questions))
	for key, question := range questions.Questions {
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
}
