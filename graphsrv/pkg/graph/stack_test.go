package graph

import (
	"reflect"
	"testing"
)

func TestStack_Remove(t *testing.T) {
	type args struct {
		ID       []string
		IDRemove string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "middle",
			args: args{
				ID:       []string{"123", "345", "567", "890"},
				IDRemove: "567",
			},
			want: []string{"123", "345", "890"},
		},
		{
			name: "first",
			args: args{
				ID:       []string{"123", "345", "567", "890"},
				IDRemove: "123",
			},
			want: []string{"345", "567", "890"},
		},
		{
			name: "last",
			args: args{
				ID:       []string{"123", "345", "567", "890"},
				IDRemove: "890",
			},
			want: []string{"123", "345", "567"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for _, id := range tt.args.ID {
				s.Push(id)
			}
			s.Remove(tt.args.IDRemove)
			reflect.DeepEqual(tt.want, s.data)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type args struct {
		ID []string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantBool bool
	}{
		{
			name: "pop",
			args: args{
				ID: []string{"123", "345", "567", "890"},
			},
			want:     "890",
			wantBool: true,
		},
		{
			name: "pop",
			args: args{
				ID: []string{},
			},
			want:     "",
			wantBool: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for _, data := range tt.args.ID {
				s.Push(data)
			}
			if got, ok := s.Pop(); got != tt.want || ok != tt.wantBool {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
