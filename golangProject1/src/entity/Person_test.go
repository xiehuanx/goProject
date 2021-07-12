package entity

import (
	"reflect"
	"testing"
)

func TestNewPerson(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *Person
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPerson(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
