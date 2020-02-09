package main

import (
	"encoding/json"
	"fmt"
)

var Tree *UserHierachyTree

func main() {
	Tree = &UserHierachyTree{}
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
	Tree.SetRoles(rolesRequest)
}

func setUsers(rolesRequest []byte) {
	Tree.SetUsers(rolesRequest)
}

func getSubordinates(userID int) {
	subordinates := Tree.GetSubordinates(userID)
	subPayload, subPayloadErr := json.Marshal(subordinates)
	if subPayloadErr != nil {
		fmt.Errorf("error when encoding users, msg: %v",
			subPayloadErr)
	}
	fmt.Println(string(subPayload))
}
