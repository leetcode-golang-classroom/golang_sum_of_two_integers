package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	num1 := 10
	num2 := 5
	for idx := 0; idx < b.N; idx++ {
		getSum(num1, num2)
	}

}
func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a = 1, b = 2",
			args: args{a: 1, b: 2},
			want: 3,
		},
		{
			name: "a = 2, b = 3",
			args: args{a: 2, b: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
