package calculator

import (
	"errors"
	"math"
)

type Option int

const (
	Call Option = iota
	Put
)

var (
	ErrInvalidStrike       = errors.New("calculator: strike should be >= 0")
	ErrInvalidStoke        = errors.New("calculator: stoke should be >= 0")
	ErrInvalidPremium      = errors.New("calculator: premium should be >= 0")
	ErrInvalidInterestRate = errors.New("calculator: interest rate should be >= 0 and <= 1")
	ErrInvalidVolatility   = errors.New("calculator: volatility should be >= 0 and <= 1")
	ErrInvalidTimeToExpire = errors.New("calculator: time to expire should be >= 0")
	ErrInvalidDividend     = errors.New("calculator: dividend should be >= 0 and <= 1")
)

type BlackScholesModel struct {
	Option
	Strike       float64
	Stock        float64
	InterestRate float64
	Volatility   float64
	TimeToExpire float64
	Dividend     float64
}

func (bsm BlackScholesModel) Price() (float64, error) {
	if bsm.Strike < 0 {
		return 0, ErrInvalidStrike
	}
	if bsm.Stock < 0 {
		return 0, ErrInvalidStoke
	}
	if bsm.InterestRate < 0 || bsm.InterestRate > 1 {
		return 0, ErrInvalidInterestRate
	}
	if bsm.Volatility < 0 || bsm.Volatility > 1 {
		return 0, ErrInvalidVolatility
	}
	if bsm.TimeToExpire < 0 {
		return 0, ErrInvalidTimeToExpire
	}
	if bsm.Dividend < 0 || bsm.Dividend > 1 {
		return 0, ErrInvalidDividend
	}
	d1 := math.Log(bsm.Stock/bsm.Strike) + bsm.InterestRate - bsm.Dividend + math.Pow(bsm.Volatility, 2)/2*bsm.TimeToExpire/bsm.Volatility*math.Sqrt(bsm.TimeToExpire)
	d2 := d1 - bsm.Volatility*math.Sqrt(bsm.TimeToExpire)
	if bsm.Option == Call {
		return bsm.Stock*math.Pow(math.E, -bsm.Dividend*bsm.TimeToExpire)*normDist(d1) - bsm.Strike*math.Pow(math.E, -bsm.InterestRate*bsm.TimeToExpire)*normDist(d2), nil
	}
	return bsm.Strike*math.Pow(math.E, -bsm.InterestRate*bsm.TimeToExpire)*normDist(-d2) - bsm.Stock*math.Pow(math.E, -bsm.Dividend*bsm.TimeToExpire)*normDist(-d1), nil
}

func BreakEvenPoint(opt Option, str, pr float64) (float64, error) {
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

func PayoffFromBuying(opt Option, str, st, pr float64) (float64, error) {
	if str < 0 {
		return 0, ErrInvalidStrike
	}
	if st < 0 {
		return 0, ErrInvalidStoke
	}
	if pr < 0 {
		return 0, ErrInvalidPremium
	}
	if opt == Call {
		return math.Max(st-str, 0) - pr, nil
	}
	return math.Max(str-st, 0) - pr, nil
}

func PayoffFromSelling(opt Option, str, st, pr float64) (float64, error) {
	if str < 0 {
		return 0, ErrInvalidStrike
	}
	if st < 0 {
		return 0, ErrInvalidStoke
	}
	if pr < 0 {
		return 0, ErrInvalidPremium
	}
	if opt == Call {
		return pr - math.Max(st-str, 0), nil
	}
	return pr - math.Max(str-st, 0), nil
}

func normDist(z float64) float64 {
	t := 1 / (1 + 0.2316419*math.Abs(z))
	t2 := math.Pow(t, 2)
	y := t * (0.319381530 - 0.356563782*t + (1.781477937-1.821255978*t+1.330274429*t2)*t2)

	if z > 0 {
		return 1 - math.Exp(-(math.Log(2*math.Pi)+math.Pow(z, 2))*0.5)*y
	}
	return math.Exp(-(math.Log(2*math.Pi)+math.Pow(-z, 2))*0.5) * y
}
