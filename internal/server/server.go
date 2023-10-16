package server

import (
	"context"
	"github.com/alexilallas/quiz/internal/core/port"
	"github.com/alexilallas/quiz/internal/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
	grpc.UnimplementedQuizServer
	validator Validator
	usecase   port.QuizUseCase
}

func (s Server) ListQuestions(context.Context, *empty.Empty) (*grpc.Questions, error) {
	return s.usecase.ListQuestions().ToDTO(), nil
}

func (s Server) RegisterAnswers(_ context.Context, a *grpc.Answer) (*grpc.QuizResponse, error) {
	if err := s.validator.Request(a); err != nil {
		return &grpc.QuizResponse{}, err
	}
	response := s.usecase.RegisterAnswers(a.Answer)
	return response.ToDTO(), nil
}

func ProvideServer(validator Validator, usecase port.QuizUseCase) grpc.QuizServer {
	return Server{
		validator: validator,
		usecase:   usecase,
	}
}
