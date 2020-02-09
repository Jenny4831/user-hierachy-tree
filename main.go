package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var userHierachyTree *UserHierachyTree

var SampleRolesRequest = []byte(`[
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

var SampleUsersRequest = []byte(`[
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
	}]`)

func setRoles(rolesRequest []byte) error {
	var roles []Role
	err := json.Unmarshal(rolesRequest, &roles)
	if err != nil {
		return fmt.Errorf("fail to unmarshal roles, errMsg: %v", err)
	}
	setRolesErr := userHierachyTree.SetRoles(roles)
	if setRolesErr != nil {
		return setRolesErr
	}
	return nil
}

func setUsers(usersRequest []byte) error {
	var users []User
	err := json.Unmarshal(usersRequest, &users)
	if err != nil {
		return fmt.Errorf("fail to unmarshal users, errMsg: %v", err)
	}
	setUsersErr := userHierachyTree.SetUsers(users)
	if setUsersErr != nil {
		return setUsersErr
	}
	return nil
}

func getSubordinates(userID int) ([]byte, error) {
	subordinates := userHierachyTree.GetSubordinates(userID)
	subPayload, subPayloadErr := json.Marshal(subordinates)
	if subPayloadErr != nil {
		return nil, fmt.Errorf("error when encoding users, msg: %v",
			subPayloadErr)
	}
	if subPayload != nil {
		fmt.Println(string(subPayload))
	}
	return subPayload, nil
}

func main() {
	//initialise tree
	userHierachyTree = &UserHierachyTree{}
	rolesRequest := SampleRolesRequest
	usersRequest := SampleUsersRequest

	setRolesErr := setRoles(rolesRequest)
	if setRolesErr != nil {
		fmt.Printf("error setting roles, msg: %v", setRolesErr.Error())
		os.Exit(1)
	}

	setUsersErr := setUsers(usersRequest)
	if setUsersErr != nil {
		fmt.Printf("error setting users, msg: %v", setUsersErr.Error())
		os.Exit(1)
	}

	//sample roles
	fmt.Printf("Sample roles\n:%s\n", string(SampleRolesRequest))

	//sample roles
	fmt.Printf("Sample users\n:%s\n", string(SampleUsersRequest))

	fmt.Println("getSubordinates(3):")
	_, subordinatesErr := getSubordinates(3)
	if subordinatesErr != nil {
		fmt.Printf("error getting subordinates, msg: %v", subordinatesErr.Error())
		os.Exit(1)
	}

	fmt.Println("getSubordinates(1):")
	_, subsErr := getSubordinates(1)
	if subsErr != nil {
		fmt.Printf("error getting subordinates, msg: %v", subsErr.Error())
		os.Exit(1)
	}
}
