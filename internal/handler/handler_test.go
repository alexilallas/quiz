package handler

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	pb "github.com/alexilallas/quiz/internal/grpc"
	"github.com/alexilallas/quiz/internal/grpc/mocks"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var testQuestions = pb.Questions{
	Questions: []*pb.Question{
		{
			Description: "Choose Option B",
			Options: map[string]*pb.Option{
				"A": {
					Description: "Description 1",
				},
				"B": {
					Description: "Description 2",
					IsCorrect:   true,
				},
			},
		},
		{
			Description: "Choose option A",
			Options: map[string]*pb.Option{
				"A": {
					Description: "Description 1",
					IsCorrect:   true,
				},
				"B": {
					Description: "Description 2",
				},
			},
		},
	}}

func TestQuizHandler(t *testing.T) {
	var (
		ctx    = context.Background()
		client = new(mocks.QuizClient)
		s      = bufio.NewScanner(strings.NewReader("A B"))
	)

	s.Split(bufio.ScanWords)
	h := ProvideHandler(s)

	client.On("ListQuestions", ctx, mock.MatchedBy(func(*empty.Empty) bool { return true })).Return(&pb.Questions{}, errors.New("server error")).Once()
	t.Run("Should return an error when is not possible to get questions", func(t *testing.T) {
		err := h.QuizHandler(ctx, client)
		assert.NotNil(t, err)
	})

	client.On("ListQuestions", ctx, mock.MatchedBy(func(*empty.Empty) bool { return true })).Return(&testQuestions, nil)
	client.On("RegisterAnswers", ctx, &pb.Answer{Answer: []string{"A", "B"}}).Return(&pb.QuizResponse{}, errors.New("failed to register answers")).Once()
	t.Run("Should return error when fail to register answers", func(t *testing.T) {
		err := h.QuizHandler(ctx, client)
		assert.NotNil(t, err)
	})

	client.On("RegisterAnswers", ctx, &pb.Answer{Answer: []string{"A", "B"}}).Return(&pb.QuizResponse{}, nil)
	t.Run("Should return error when scanning ", func(t *testing.T) {
		err := h.QuizHandler(ctx, client)
		assert.NotNil(t, err)
		assert.Equal(t, ErrorScanAnswer{}.Error(), err.Error())
	})

	t.Run("Should make an successful request ", func(t *testing.T) {
		s = bufio.NewScanner(strings.NewReader("A B"))
		s.Split(bufio.ScanWords)
		err := ProvideHandler(s).QuizHandler(ctx, client)
		assert.Nil(t, err)
	})
}

func generateText(n int) string {
	var str string
	for i := 0; i < n; i++ {
		str += strconv.Itoa(rand.Int()) + " "
	}
	return str
}

func BenchmarkBufferIo(b *testing.B) {
	str := generateText(b.N)
	reader := strings.NewReader(str)
	s := bufio.NewScanner(reader)
	s.Split(bufio.ScanWords)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !s.Scan() {
			b.Fatal("s.Scan is false")
		}
	}
}

func BenchmarkScan(b *testing.B) {
	str := generateText(b.N)
	reader := strings.NewReader(str)
	var val string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := fmt.Fscan(reader, &val); err != nil {
			b.Fatal(err)
		}
	}
}
