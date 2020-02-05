package main

import (
	"reflect"
	"testing"
)

func TestGetSubOrdinates(t *testing.T) {

	tests := []struct {
		name   string
		userID int
		want   []User
	}{
		{
			name:   "sample a, success",
			userID: 3,
			want: []User{
				{Id: 2, Name: "Emily Employee", Role: 4},
				{Id: 5, Name: "Steve Trainer", Role: 5},
			},
		},
		{
			name:   "sample b, success",
			userID: 1,
			want: []User{
				{Id: 2, Name: "Emily Employee", Role: 4},
				{Id: 3, Name: "Sam Supervisor", Role: 3},
				{Id: 4, Name: "Mary Manager", Role: 2},
				{Id: 5, Name: "Steve Trainer", Role: 5},
			},
		},
		{
			name:   "no subordinates",
			userID: 5,
			want:   []User{},
		},
		{
			name:   "user record does not exist",
			userID: 1000,
			want:   []User{},
		},
		{
			name:   "user record does not exist",
			userID: -1,
			want:   []User{},
		},
		// {
		// 	name:   "role does not exist for given user",
		// 	userID: ,
		// 	want:   []User{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSubOrdinates(tt.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubOrdinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
