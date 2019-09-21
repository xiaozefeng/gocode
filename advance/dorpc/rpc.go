package dorpc

import "errors"

type Service struct {
}

type ServiceArgs struct {
	A, B int
}

func (s Service) Div(args ServiceArgs, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
