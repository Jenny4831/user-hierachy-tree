package main

import (
	"fmt"
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

func TestSetRoles(t *testing.T) {
	type fields struct {
		RolesMap map[int]Role
		Tree     *UserHierachyTree
	}
	tests := []struct {
		name   string
		fields fields
		data   []byte
	}{
		{
			name: "unmarshall success",
			data: []byte(`[
				{
				"Id": 1,
				"Name": "System Administrator",
				"Parent": 0
				},
				{
				"Id": 2,
				"Name": "Location Manager",
				"Parent": 1
				},
				{
				"Id": 3,
				"Name": "Supervisor",
				"Parent": 2
				},
				{
				"Id": 4,
				"Name": "Employee",
				"Parent": 3
				},
				{
				"Id": 5,
				"Name": "Trainer",
				"Parent": 3
				}
			 ]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userHierachy := &UserHierachy{
				RolesMap: tt.fields.RolesMap,
				Tree:     tt.fields.Tree,
			}
			userHierachy.SetRoles(tt.data)
		})
	}
}

func TestUserHierachy_SetUsers(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		RolesMap map[int]Role
		Tree     *UserHierachyTree
	}{
		// TODO: Add test cases.
		{name: "set users",
			data: []byte(`[
			{
			"Id": 1,
			"Name": "Adam Admin",
			"Role": 1
			},
			{
			"Id": 2,
			"Name": "Emily Employee",
			"Role": 4
			},
			{
			"Id": 3,
			"Name": "Sam Supervisor",
			"Role": 3
			},
			{
			"Id": 4,
			"Name": "Mary Manager",
			"Role": 2
			},
			{"Id": 5,
			"Name": "Steve Trainer",
			"Role": 5
			}]`,
			),
			RolesMap: map[int]Role{
				1: {
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				2: {
					Id:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
				3: {
					Id:     3,
					Name:   "Supervisor",
					Parent: 2,
				},
				4: {
					Id:     4,
					Name:   "Employee",
					Parent: 3,
				},
				5: {
					Id:     5,
					Name:   "Trainer",
					Parent: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userHierachy := &UserHierachy{
				RolesMap: tt.RolesMap,
				Tree:     &UserHierachyTree{},
			}
			userHierachy.SetUsers(tt.data)
			fmt.Println(userHierachy.Tree)
		})
	}
}
