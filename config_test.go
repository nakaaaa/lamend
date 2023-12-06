package lamend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalConfigYAML(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "success unmarshal yml",
			args: args{
				b: []byte("function_name: test\nimage_uri: test_uri"),
			},
			want: &Config{
				FunctionName: "test",
				ImageURI:     "test_uri",
			},
		},
		{
			name: "empty function_name",
			args: args{
				b: []byte("function_name: \nimage_uri: test_uri"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty image_uri",
			args: args{
				b: []byte("function_name: test\nimage_uri: "),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalConfigYAML(tt.args.b)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
