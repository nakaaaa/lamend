package lamend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		err     error
		wantErr bool
	}{
		{
			name: "success read yml file",
			args: args{
				path: "testdata/test.yml",
			},
			wantErr: false,
		},
		{
			name: "success read yaml file",
			args: args{
				path: "testdata/test.yaml",
			},
			wantErr: false,
		},
		{
			name: "notfound file",
			args: args{
				path: "testdata/notfound.yml",
			},
			err:     errUnsupportedFileFormat,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read(tt.args.path)
			assert.Equal(t, tt.wantErr, err != nil)
			if err != nil {
				assert.Error(t, tt.err, err)
			}
		})
	}
}
