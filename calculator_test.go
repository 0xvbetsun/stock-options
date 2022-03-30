package calculator

import (
	"math"
	"testing"
)

func TestBreakEvenPoint(t *testing.T) {
	type args struct {
		opt Option
		str float64
		pr  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "option with incorrect 'strike'", args: args{str: -50}, wantErr: true},
		{name: "'option with incorrect 'premium'", args: args{pr: -10}, wantErr: true},
		{name: "'call' with 'premium'", args: args{opt: Call, str: 50, pr: 10}, want: 60},
		{name: "'call' without 'premium'", args: args{opt: Call, str: 50}, want: 50},
		{name: "'put' with 'premium'", args: args{opt: Put, str: 50, pr: 10}, want: 40},
		{name: "'put' without 'premium'", args: args{opt: Put, str: 50}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BreakEvenPoint(tt.args.opt, tt.args.str, tt.args.pr)
			if (err != nil) != tt.wantErr {
				t.Errorf("BreakEvenPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BreakEvenPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayoffFromBuying(t *testing.T) {
	type args struct {
		opt Option
		str float64
		st  float64
		pr  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "incorrect strike", args: args{str: -50}, wantErr: true},
		{name: "incorrect stoke", args: args{st: -50}, wantErr: true},
		{name: "incorrect premium", args: args{pr: -10}, wantErr: true},
		{name: "profit from Call when 'strike' without 'premium' is less than 'stock'", args: args{opt: Call, str: 50, st: 70}, want: 20},
		{name: "profit from Call when 'strike' with 'premium' is less than 'stock'", args: args{opt: Call, str: 50, st: 70, pr: 10}, want: 10},
		{name: "no profit from Call when 'strike' without 'premium' is more than 'stock'", args: args{opt: Call, str: 50, st: 40}, want: 0},
		{name: "lose from Call when 'strike' with 'premium' is more than 'stock'", args: args{opt: Call, str: 50, st: 10, pr: 10}, want: -10},
		{name: "profit from Put when 'strike' without 'premium' is more than 'stock'", args: args{opt: Put, str: 50, st: 40}, want: 10},
		{name: "profit from Put when 'strike' with 'premium' is more than 'stock'", args: args{opt: Put, str: 50, st: 30, pr: 10}, want: 10},
		{name: "no profit from Put when 'strike' without 'premium' is less than 'stock'", args: args{opt: Put, str: 50, st: 60}, want: 0},
		{name: "lose from Put when 'strike' with 'premium' is less than 'stock'", args: args{opt: Put, str: 50, st: 70, pr: 10}, want: -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PayoffFromBuying(tt.args.opt, tt.args.str, tt.args.st, tt.args.pr)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayoffFromBuying() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PayoffFromBuying() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayoffFromSelling(t *testing.T) {
	type args struct {
		opt Option
		str float64
		st  float64
		pr  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "incorrect strike", args: args{str: -50}, wantErr: true},
		{name: "incorrect stoke", args: args{st: -50}, wantErr: true},
		{name: "incorrect premium", args: args{pr: -10}, wantErr: true},
		{name: "profit from Call when 'strike' with 'premium' is more than 'stock'", args: args{opt: Call, str: 50, st: 40, pr: 10}, want: 10},
		{name: "no profit from Call when 'strike' without 'premium' is more than 'stock'", args: args{opt: Call, str: 50, st: 40}, want: 0},
		{name: "lose from Call when 'strike' without 'premium' is more than 'stock'", args: args{opt: Call, str: 50, st: 60}, want: -10},
		{name: "lose from Call when 'strike' with 'premium' is less than 'stock'", args: args{opt: Call, str: 50, st: 70, pr: 10}, want: -10},
		{name: "infinity lose from Call when 'stock' gets infinity raise", args: args{opt: Call, str: 50, st: math.Inf(1)}, want: math.Inf(-1)},

		{name: "profit from Put when 'strike' with 'premium' is less than 'stock'", args: args{opt: Put, str: 50, st: 70, pr: 10}, want: 10},
		{name: "no profit from Put when 'strike' without 'premium' is more than 'stock'", args: args{opt: Put, str: 50, st: 60}, want: 0},
		{name: "lose from Put when 'strike' without 'premium' is more than 'stock'", args: args{opt: Put, str: 50, st: 40}, want: -10},
		{name: "lose from Put when 'strike' with 'premium' is more than 'stock'", args: args{opt: Put, str: 50, st: 30, pr: 10}, want: -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PayoffFromSelling(tt.args.opt, tt.args.str, tt.args.st, tt.args.pr)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayoffFromSelling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PayoffFromSelling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlackScholesModel_Price(t *testing.T) {
	type fields struct {
		Option       Option
		Strike       float64
		Stock        float64
		InterestRate float64
		Volatility   float64
		TimeToExpire float64
		Dividend     float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{name: "invalid strike", fields: fields{Strike: -1}, wantErr: true},
		{name: "invalid stoke", fields: fields{Stock: -1}, wantErr: true},
		{name: "invalid interest rate", fields: fields{InterestRate: 2}, wantErr: true},
		{name: "invalid volatility", fields: fields{Volatility: 2}, wantErr: true},
		{name: "invalid time to expire", fields: fields{TimeToExpire: -1}, wantErr: true},
		{name: "invalid dividend", fields: fields{Dividend: 2}, wantErr: true},
		{name: "call with dividends", fields: fields{Option: Call, Strike: 58, Stock: 60, InterestRate: 0.035, Volatility: 0.2, TimeToExpire: 0.5, Dividend: 0.0125}, want: 4.624545765201692},
		{name: "call without dividends", fields: fields{Option: Call, Strike: 58, Stock: 60, InterestRate: 0.035, Volatility: 0.2, TimeToExpire: 0.5}, want: 4.838950566319454},
		{name: "put with dividends", fields: fields{Option: Put, Strike: 58, Stock: 60, InterestRate: 0.035, Volatility: 0.2, TimeToExpire: 0.5, Dividend: 0.0125}, want: 1.9922059963722596},
		{name: "put without dividends", fields: fields{Option: Put, Strike: 58, Stock: 60, InterestRate: 0.035, Volatility: 0.2, TimeToExpire: 0.5}, want: 1.8327802348937041},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bsm := BlackScholesModel{
				Option:       tt.fields.Option,
				Strike:       tt.fields.Strike,
				Stock:        tt.fields.Stock,
				InterestRate: tt.fields.InterestRate,
				Volatility:   tt.fields.Volatility,
				TimeToExpire: tt.fields.TimeToExpire,
				Dividend:     tt.fields.Dividend,
			}
			got, err := bsm.Price()
			if (err != nil) != tt.wantErr {
				t.Errorf("BlackScholesModel.Price() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BlackScholesModel.Price() = %v, want %v", got, tt.want)
			}
		})
	}
}
