package jsonparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	type SimpleTest struct {
		Key string `json:"key"`
	}

	type args struct {
		str string
		v   SimpleTest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				str: "{\"key\":\"value\"}\n",
				v:   SimpleTest{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parse(tt.args.str, &tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, "value", tt.args.v.Key)
		})
	}
}
