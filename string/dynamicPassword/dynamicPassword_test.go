package dynamicpassword

import "testing"

func Test_dynamicPassword(t *testing.T) {
	type args struct {
		password string
		target   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "用例一",
			args: args{
				password: "s3cur1tyC0d3",
				target:   4,
			},
			want: "r1tyC0d3s3cu",
		},
		{
			name: "用例二",
			args: args{
				password: "lrloseumgh",
				target:   6,
			},
			want: "umghlrlose",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dynamicPassword(tt.args.password, tt.args.target); got != tt.want {
				t.Errorf("dynamicPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
