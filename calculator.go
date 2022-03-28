package calculator

import "errors"

type Option int

const (
	Call Option = iota
	Put
)

var (
	ErrInvalidStrike  = errors.New("calculator: strike should be >= 0")
	ErrInvalidPremium = errors.New("calculator: premium should be >= 0")
)

func BreakEvenPoint(opt Option, str, pr float32) (float32, error) {
	if str < 0 {
		return 0, ErrInvalidStrike
	}
	if pr < 0 {
		return 0, ErrInvalidPremium
	}

	if opt == Call {
		return str + pr, nil
	}
	return str - pr, nil
}
