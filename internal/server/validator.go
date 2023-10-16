package server

import (
	"errors"
	pb "github.com/alexilallas/quiz/internal/grpc"
)

type Validator struct{}

func (Validator) Request(r *pb.Answer) error {
	if r == nil || len(r.Answer) == 0 {
		return ErrorAnswersIsNil{}
	}

	var err error
	for _, v := range r.Answer {
		if v != "A" && v != "B" && v != "C" {
			err = errors.Join(err, ErrorInvalidOption(v))
		}
	}
	return err
}
