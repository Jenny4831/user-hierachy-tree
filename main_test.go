package main

import (
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name             string
		roles            []byte
		wantRolesErr     bool
		users            []byte
		wantUsersErr     bool
		userId           int
		wantSubordinates []byte
	}{
		{
			name:         "test full interaction, random user id",
			roles:        SampleRolesRequest,
			wantRolesErr: false,
			users:        SampleUsersRequest,
			wantUsersErr: false,
			userId:       3,
			wantSubordinates: []byte(`[{"Id":2,"Name":"Emily Employee","Role":4},` +
				`{"Id":5,"Name":"Steve Trainer","Role":5}]`),
		},
		{
			name:         "test full interaction, user id in root",
			roles:        SampleRolesRequest,
			wantRolesErr: false,
			users:        SampleUsersRequest,
			wantUsersErr: false,
			userId:       1,
			wantSubordinates: []byte(`[{"Id":4,"Name":"Mary Manager","Role":2},{"Id":3,"Name":"Sam Supervisor","Role":3},` +
				`{"Id":2,"Name":"Emily Employee","Role":4},{"Id":5,"Name":"Steve Trainer","Role":5}]`),
		},
		{
			name:             "test full interaction, invalid roles request",
			roles:            []byte(``),
			wantRolesErr:     true,
			users:            SampleUsersRequest,
			wantUsersErr:     false,
			userId:           1,
			wantSubordinates: []byte(`[]`),
		},
		{
			name:             "test full interaction, user has no subordinates",
			roles:            SampleRolesRequest,
			wantRolesErr:     false,
			users:            SampleUsersRequest,
			wantUsersErr:     false,
			userId:           5,
			wantSubordinates: []byte(`[]`),
		},
		{
			name:             "test full interaction, no users set",
			roles:            SampleRolesRequest,
			wantRolesErr:     false,
			users:            []byte(``),
			wantUsersErr:     false,
			userId:           5,
			wantSubordinates: []byte(`[]`),
		},
		{
			name:             "test full interaction, invalid users request",
			roles:            SampleRolesRequest,
			wantRolesErr:     false,
			users:            []byte(`["randomObject": "random",]`),
			wantUsersErr:     true,
			userId:           5,
			wantSubordinates: []byte(`[]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userHierachyTree = &UserHierachyTree{}
			gotRolesErr := setRoles(tt.roles)
			if gotRolesErr == nil && tt.wantRolesErr {
				t.Errorf("setRoles = %v, wantRolesErr %v",
					gotRolesErr, tt.wantRolesErr)
			}
			if gotRolesErr != nil {
				//exit when userHierachytree failed to generate
				return
			}
			setUsersErr := setUsers(tt.users)
			if setUsersErr == nil && tt.wantUsersErr {
				t.Errorf("setUsers = %v, wantUsersErr %v",
					setUsersErr, tt.wantUsersErr)
			}
			got, _ := getSubordinates(tt.userId)
			if !strings.EqualFold(string(got), string(tt.wantSubordinates)) {
				t.Errorf("getSubordinates = %v, want %v",
					string(got), string(tt.wantSubordinates))
			}
		})
	}
}
