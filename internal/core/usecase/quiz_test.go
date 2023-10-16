package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuiz_ListQuestions(t *testing.T) {
	var usecase = ProvideQuizUseCase()

	t.Run("Should return an list of questions", func(t *testing.T) {
		actual := usecase.ListQuestions()
		assert.Equal(t, questions, actual)
	})
}

func TestQuiz_RegisterAnswers(t *testing.T) {
	var usecase = ProvideQuizUseCase()

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
			actual := usecase.RegisterAnswers(tt.answer)
			assert.Equal(t, tt.expected.correct, actual.Correct)
			assert.Equal(t, tt.expected.totalQuestions, actual.TotalQuestions)
			assert.Equal(t, tt.expected.percentage, actual.Percentage)
		})
	}
}
