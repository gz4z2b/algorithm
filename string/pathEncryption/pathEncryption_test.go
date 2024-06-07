package pathencryption

import "testing"

func Test_pathEncryption(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "用例一",
			args: args{path: "a.aef.qerf.bb"},
			want: "a aef qerf bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathEncryption(tt.args.path); got != tt.want {
				t.Errorf("pathEncryption() = %v, want %v", got, tt.want)
			}
		})
	}
}
