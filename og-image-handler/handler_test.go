package function

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_getOgImage(t *testing.T) {
	tests := []struct {
		name    string
		want    *bytes.Buffer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getOgImage("../template2.jpg", "Hello World")
			if (err != nil) != tt.wantErr {
				t.Errorf("getOgImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOgImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
