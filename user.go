package main

type Role struct {
	Id     int
	Name   string
	Parent int
}

type User struct {
	Id   int
	Name string
	Role int
}

func SetRoles(roles []Role) interface{} {
	return nil
}

func SetUsers(users []User) interface{} {
	return nil
}

func GetSubOrdinates(userID int) []User {
	var users []User
	return users
}

