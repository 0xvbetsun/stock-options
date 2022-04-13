// Package calculator provides basic constants and mathematical functions for stock options.
package calculator

import (
	"errors"
	"math"
)

type Option int

// Option types
const (
	Call Option = iota
	Put
)

// Validation errors
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
	Option               // option type (call or put)
	Strike       float64 // strike price ($$$ per share)
	Stock        float64 // underlying price ($$$ per share)
	InterestRate float64 // continuously compounded risk-free interest rate (% p.a.)
	Volatility   float64 // volatility (% p.a.)
	TimeToExpire float64 // time to expiration (% of year)
	Dividend     float64 // continuously compounded dividend yield (% p.a.)
}

// Price calculates by solving the equation for the corresponding terminal and boundary conditions
// for more details see https://en.wikipedia.org/wiki/Black-Scholes_model
func (bsm *BlackScholesModel) Price() (float64, error) {
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

// BreakEvenPoint calculates price in the underlying asset at which exercise/dispose
// of the contract without incurring a loss
//
// It takes strike price (str) and premium (pr) in $$$ per share and returns break-even point
// regarding to the option type (opt)
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

// PayoffFromBuying calculates current market price of an option ($$$ per share) from buying
//
// It takes strike price (str), stock (st) and premium (pr) in $$$ per share and returns profit/lose
// from buying options regarding to the option  type (opt)
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

// PayoffFromSelling calculates current market price of an option ($$$ per share) from selling
//
// It takes strike price (str), stock (st) and premium (pr) in $$$ per share and returns profit/lose
// from selling options regarding to the option type (opt)
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
