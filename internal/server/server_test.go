package server

import (
	"context"
	"github.com/alexilallas/quiz/internal/core/entity"
	"github.com/alexilallas/quiz/internal/core/port/mocks"
	"github.com/alexilallas/quiz/internal/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServer_ListQuestions(t *testing.T) {
	var (
		ctx     = context.Background()
		usecase = new(mocks.QuizUseCase)
		server  = ProvideServer(Validator{}, usecase)
	)

	usecase.On("ListQuestions").Return(entity.Questions{})

	t.Run("Should return an list of questions", func(t *testing.T) {
		_, err := server.ListQuestions(ctx, nil)
		assert.Nil(t, err)
	})
}

func TestServer_RegisterAnswers(t *testing.T) {
	var (
		ctx     = context.Background()
		usecase = new(mocks.QuizUseCase)
		server  = ProvideServer(Validator{}, usecase)
	)

	usecase.On("ListQuestions").Return(entity.Questions{})

	t.Run("Should fail validation when answer is empty", func(t *testing.T) {
		actual, err := server.RegisterAnswers(ctx, &grpc.Answer{})
		assert.NotNil(t, err)
		assert.Equal(t, ErrorAnswersIsNil{}.Error(), err.Error())
		assert.Zero(t, actual.Correct)
		assert.Zero(t, actual.TotalQuestions)
		assert.Zero(t, actual.Percentage)
	})

	t.Run("Should fail validation when option doesn't exist", func(t *testing.T) {
		actual, err := server.RegisterAnswers(ctx, &grpc.Answer{Answer: []string{"Z"}})
		assert.NotNil(t, err)
		assert.Equal(t, ErrorInvalidOption("Z").Error(), err.Error())
		assert.Zero(t, actual.Correct)
		assert.Zero(t, actual.TotalQuestions)
		assert.Zero(t, actual.Percentage)
	})

	t.Run("Should register successfully", func(t *testing.T) {
		usecase.On("RegisterAnswers", mock.MatchedBy(func(entity.Answer) bool { return true })).Return(entity.Response{})
		_, err := server.RegisterAnswers(ctx, &grpc.Answer{Answer: []string{"A", "B", "C"}})
		assert.Nil(t, err)
	})

}
