package product_service

import "testing"

func TestIsPalindrom(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive",
			args: args{"asddsa"},
			want: true,
		},
		{
			name: "negative",
			args: args{"ovi hzjyvui"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrom(tt.args.query); got != tt.want {
				t.Errorf("IsPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
