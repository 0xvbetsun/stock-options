package calculator

import "testing"

func TestBreakEvenPoint(t *testing.T) {
	type args struct {
		opt Option
		str float32
		pr  float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{name: "Call option with premium", args: args{opt: Call, str: 50, pr: 10}, want: 60},
		{name: "Call option without premium", args: args{opt: Call, str: 50}, want: 50},
		{name: "Call option with incorrect premium", args: args{opt: Call, str: 50, pr: -10}, wantErr: true},
		{name: "Call option with incorrect strike", args: args{opt: Call, str: -50, pr: 10}, wantErr: true},
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
