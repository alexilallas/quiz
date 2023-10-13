package handler

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/alexilallas/quiz/internal/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"strings"
)

const separator = "-------------------------------------"

type Handler struct {
	scanner *bufio.Scanner
}

func (h Handler) QuizHandler(ctx context.Context, client pb.QuizClient) error {
	questions, err := client.ListQuestions(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	var answer pb.Answer

	if err = h.getAnswers(questions, &answer); err != nil {
		return err
	}

	r, err := client.RegisterAnswers(ctx, &answer)
	if err != nil {
		return err
	}

	fmt.Printf("You got %0.f out of %.0f.\nYou were better than %.2f%% of all quizzers\n", r.Correct, r.TotalQuestions, r.Percentage)
	return nil
}

func (h Handler) getAnswers(questions *pb.Questions, answer *pb.Answer) error {
	fmt.Println(separator)
	for n, question := range questions.Questions {
		fmt.Printf("Question %d: %s\n", n+1, question.Description)
		for key, option := range question.Options {
			fmt.Printf("%s) %s\n", key, option.Description)
		}

		if !h.scanner.Scan() {
			return ErrorScanAnswer{}
		}
		answer.Answer = append(answer.Answer, strings.ToUpper(h.scanner.Text()))
		fmt.Println(separator)
	}
	return nil
}

func ProvideHandler(scanner *bufio.Scanner) Handler {
	return Handler{
		scanner: scanner,
	}
}
