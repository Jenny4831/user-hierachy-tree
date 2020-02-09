package main

import (
	"encoding/json"
	"fmt"
)

var userHierachyTree *UserHierachyTree

func main() {

	userHierachyTree = &UserHierachyTree{}
	rolesRequest := []byte(`[
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
	 ]`)

	usersRequest := []byte(`[
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
	)

	setRoles(rolesRequest)
	setUsers(usersRequest)
	getSubordinates(3)
}

func setRoles(rolesRequest []byte) {
	var roles []Role
	err := json.Unmarshal(rolesRequest, &roles)
	if err != nil {
		fmt.Printf("fail to unmarshal roles, errMsg: %v", err)
		return
	}
	userHierachyTree.SetRoles(roles)
}

func setUsers(usersRequest []byte) {
	var users []User
	err := json.Unmarshal(usersRequest, &users)
	if err != nil {
		fmt.Printf("fail to unmarshal users, errMsg: %v", err)
		return
	}
	userHierachyTree.SetUsers(users)
}

func getSubordinates(userID int) {
	subordinates := userHierachyTree.GetSubordinates(userID)
	subPayload, subPayloadErr := json.Marshal(subordinates)
	if subPayloadErr != nil {
		fmt.Errorf("error when encoding users, msg: %v",
			subPayloadErr)
	}
	fmt.Println(string(subPayload))
}
