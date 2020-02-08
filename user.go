package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type IUserHierachyService interface {
	SetUsers(data []byte) []User
	SetRoles(data []byte) []Role
}
type UserHierachy struct {
	RolesMap map[int]Role
	Tree     *UserHierachyTree
}
type Role struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Parent int    `json:"Parent"`
}

type User struct {
	Id   int
	Name string
	Role int
}

// each node stores a map of users using user id as key,
// User as value
// using a map since there can be multiple users with the same role
type UserNode struct {
	RolesMap    map[int]*Role
	Users       map[int]*User
	Subordinate *UserNode
}

// store hierachy of users
type UserHierachyTree struct {
	Root *UserNode
}

func NewUserHierachy() *UserHierachy {
	return &UserHierachy{
		RolesMap: make(map[int]Role),
		Tree:     &UserHierachyTree{},
	}
}

// var RolesMap = make(map[int]Role)

// func BFS(roleTree *RoleNode, role int) []User {
// 	queue := []*RoleNode{}
// 	queue = append(queue, roleTree)
// 	result := []User{}
// 	return BFSUtil(queue, result, role)
// 	return []User{}
// }

// var root = 0

// func BFSUtil(queue []*RoleNode, result []User, role int) []User {
// 	if len(queue) == 0 {
// 		return result
// 	}
// 	if queue[0].Left != nil && queue[0].Left.Value.Parent < role {

// 	}
// 	return BFSUtil(queue[1:], result, )
// }

//idea O(n)
// linked list
//linklist node {val int, Parent, list of users}
// 3 -> 2 -> 1 -> 0
// if node.val == user.Role
// return users
// users = append(users, node.users...)
//node = node.parent

//second  idea

//users in map
// user, found := map[id]User{etc}
//
// roles sorted create binary tree
//    0
//  1  4
// 2 3 5 5

// sample 0
//       1(0)
//      2(1)
//    3(2)
//  4(3)   5(3)
// func(n *Node) FindSubordinatesForRole(role int)
// users
//     if n == nil return
//     if role > n.left.role
//        users = append(users, n.left.users)
//        return FindSubordinatesForRole(n.left.role)
//     return FindSubordinatesForRole(n.right.role)
// getSubordinates()
//  FindSubordinatesForRole(user.role)
//
//

// unmarshall data and store roles in public role map,
// key -> ID, value -> role
func (userHierachy *UserHierachy) SetRoles(data []byte) {
	var roles []Role
	err := json.Unmarshal(data, &roles)
	if err != nil {
		fmt.Println("fail to unmarshal roles")
	}

	rolesMap := make(map[int]Role, len(roles))
	if len(roles) == 0 {
		fmt.Println("Roles list is empty")
		return
	}
	for idx := range roles {
		role := roles[idx]
		rolesMap[role.Id] = role
	}
	userHierachy.RolesMap = rolesMap
}

// var UserHierachyTree = (root *UserNode
// 	root *UserNode
// )
//sort users by role helpes constructing `UserHierachyTree`
func (userHierachy *UserHierachy) SetUsers(data []byte) {
	var users []User
	err := json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("fail to unmarshal roles")
	}
	sortUsersByRole(users)
	for idx := range users {
		user := users[idx]
		userRole := userHierachy.RolesMap[user.Role]
		userHierachy.Tree.add(user, userRole)
	}
}

func (tree *UserHierachyTree) add(user User, userRole Role) {
	if tree.Root == nil && userRole.Parent == 0 {
		userMap := make(map[int]*User)
		userMap[user.Id] = &user
		rolesMap := make(map[int]*Role)
		rolesMap[userRole.Id] = &userRole
		userNode := UserNode{Users: userMap, RolesMap: rolesMap}
		tree.Root = &userNode
	} else {
		tree.Root.insert(user, userRole)
	}
}

// if inserting user's role parent equals to userNode's role
// add inserting user as Subordinate of user node
func (userNode *UserNode) insert(user User, role Role) {

	if userNode.RolesMap[role.Parent] != nil {
		if userNode.Subordinate == nil {
			userMap := make(map[int]*User)
			userMap[user.Id] = &user
			rolesMap := make(map[int]*Role)
			rolesMap[role.Id] = &role
			subordinateNode := UserNode{Users: userMap, RolesMap: rolesMap}
			userNode.Subordinate = &subordinateNode
		} else if userNode.Subordinate.Users[user.Id] == nil {
			userNode.Subordinate.Users[user.Id] = &user
			if userNode.Subordinate.RolesMap[role.Id] == nil {
				userNode.Subordinate.RolesMap[role.Id] = &role
				fmt.Println(userNode.Subordinate.RolesMap)
			}
		}
	} else {
		//check if should add to usr map

		userNode.Subordinate.insert(user, role)
	}
}

func sortUsersByRole(users []User) {
	// sorteUsers := users
	sort.Slice(users,
		func(i, j int) bool {
			return users[i].Role < users[j].Role
		})
	// return users
}
func GetSubOrdinates(userID int) []User {
	var users []User
	return users
}

// questions
