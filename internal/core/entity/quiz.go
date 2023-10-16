package entity

import "github.com/alexilallas/quiz/internal/grpc"

type Questions []Question

type Question struct {
	Description string
	Options     map[string]*Option
}

func (q Questions) ToDTO() *grpc.Questions {
	var questions grpc.Questions
	for _, question := range q {
		questions.Questions = append(questions.Questions, &grpc.Question{
			Options: q.toOptions(question.Options),
		})
	}
	return &questions
}

func (q Questions) toOptions(o map[string]*Option) map[string]*grpc.Option {
	var options = make(map[string]*grpc.Option, len(o))
	for k, v := range o {
		options[k] = &grpc.Option{Description: v.Description, IsCorrect: v.IsCorrect}
	}
	return options
}

type Option struct {
	Description string
	IsCorrect   bool
}

type Answer []string

type Response struct {
	Correct        float32
	TotalQuestions float32
	Percentage     float32
}

func (r Response) ToDTO() *grpc.QuizResponse {
	return &grpc.QuizResponse{
		Correct:        r.Correct,
		TotalQuestions: r.TotalQuestions,
		Percentage:     r.Percentage,
	}
}
