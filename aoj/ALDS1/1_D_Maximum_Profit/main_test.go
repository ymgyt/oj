package main

import "github.com/google/go-cmp/cmp"

func TestMain(t *teting.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `6
5
3
1
3
4
3`,
			want: `3`,
		},
		{
			input: `3
4
3
2`,
			want: `-1`,
		},
	}

	for _, tt := range tests {
		got, want := Run(tt.input), tt.want
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("(-want +got)\n%s", diff)
		}
	}
}
