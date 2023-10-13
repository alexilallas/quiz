package server

import (
	"context"
	"github.com/alexilallas/quiz/internal/grpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_ListQuestions(t *testing.T) {
	var (
		ctx    = context.Background()
		server = Server{}
	)

	t.Run("Should return an list of questions", func(t *testing.T) {
		actual, err := server.ListQuestions(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, &questions, actual)
	})
}

func TestServer_RegisterAnswers(t *testing.T) {
	var (
		ctx    = context.Background()
		server = Server{}
	)

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

	type expected struct {
		correct        float32
		totalQuestions float32
		percentage     float32
	}
	type testsCase struct {
		name     string
		answer   []string
		expected expected
	}
	var tests = []testsCase{
		{
			name:   "Should return an successful response for first quiz",
			answer: []string{"A", "B", "C"},
			expected: expected{
				correct:        1,
				totalQuestions: 3,
				percentage:     0,
			},
		},
		{
			name:   "Should return an successful response for quizzer 2",
			answer: []string{"C", "B", "C"},
			expected: expected{
				correct:        2,
				totalQuestions: 3,
				percentage:     100,
			},
		},
		{
			name:   "Should return an successful response for quizzer 3",
			answer: []string{"C", "B", "C"},
			expected: expected{
				correct:        2,
				totalQuestions: 3,
				percentage:     50,
			},
		},
		{
			name:   "Should return an successful response for quizzer 4",
			answer: []string{"C", "B", "C"},
			expected: expected{
				correct:        2,
				totalQuestions: 3,
				percentage:     33.333336,
			},
		},
		{
			name:   "Should return an successful response for quizzer 5",
			answer: []string{"C", "B", "A"},
			expected: expected{
				correct:        3,
				totalQuestions: 3,
				percentage:     100,
			},
		},
		{
			name:   "Should return an successful response for quizzer 6",
			answer: []string{"C", "B", "A"},
			expected: expected{
				correct:        3,
				totalQuestions: 3,
				percentage:     80,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := server.RegisterAnswers(ctx, &grpc.Answer{Answer: tt.answer})
			assert.Nil(t, err)
			assert.Equal(t, tt.expected.correct, actual.Correct)
			assert.Equal(t, tt.expected.totalQuestions, actual.TotalQuestions)
			assert.Equal(t, tt.expected.percentage, actual.Percentage)
		})
	}
}
