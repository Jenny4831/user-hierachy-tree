package main

import (
	"fmt"
	"sort"
)

//Role struct
//assumption:
//Id is unique
type Role struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Parent int    `json:"Parent"`
}

//User struct
//assumption:
//Id is unique
type User struct {
	Id   int
	Name string
	Role int
}

// Each node stores
// Role, Users and Subordinates
// Users store users using map, with user id as key and user as Value
// allows more efficient check whether user is in current node's Role
// Subordinates stores list of TreeNode
// Subordinate's role parent is current TreeNode
// assumption:
// multiple roles can have the same parent
type TreeNode struct {
	Role         Role
	Users        map[int]*User
	Subordinates []*TreeNode
}

// store hierachy of Roles in tree form
type UserHierachyTree struct {
	Root *TreeNode
}

// Creates new tree node with given Role
func NewTreeNode(role Role) *TreeNode {
	return &TreeNode{
		Role:         role,
		Users:        make(map[int]*User),
		Subordinates: []*TreeNode{},
	}
}

// unmarshall data and store roles in UserHierachyTree,
// sort roles list by role.Parent, this helps building the tree in role order
// loops through sorted list and prepares tree with roles
func (tree *UserHierachyTree) SetRoles(roles []Role) error {
	if len(roles) == 0 {
		err := fmt.Errorf("list of roles is empty")
		return err
	}
	sortRolesByParent(roles)
	for idx := range roles {
		role := roles[idx]
		if tree.Root == nil {
			tree.Root = NewTreeNode(role)
		} else {
			tree.Root.insertRole(role)
		}
	}
	return nil
}

// helper function to add tree node to UserHierachyTree
// if inserting role's parent is current node,
//create new node and append to current node subordinates
// otherwise if no subordinates for given node, return
// if node has subordinates, loop through subordinates
// and insert once role's parent is found
func (treeNode *TreeNode) insertRole(role Role) {
	if treeNode.Role.Id == role.Parent {
		newNode := NewTreeNode(role)
		treeNode.Subordinates = append(treeNode.Subordinates, newNode)
	} else {
		for idx := range treeNode.Subordinates {
			subordinate := treeNode.Subordinates[idx]
			subordinate.insertRole(role)
		}
	}
}

//sort users by role helps mapping users to their give role,
//since roles are already sorted in tree
//loops through users list and updates UserHierachyTree
func (tree *UserHierachyTree) SetUsers(users []User) error {
	if len(users) == 0 {
		err := fmt.Errorf("list of users is empty")
		return err
	}
	sortUsersByRole(users)
	for idx := range users {
		user := users[idx]
		tree.Root.InsertUser(user)
	}
	return nil
}

// if found user's role in current node's role, add user to Users map
// else loop through subordinates and recursively find treenode with user's role
func (treeNode *TreeNode) InsertUser(user User) {
	if treeNode.Role.Id == user.Role {
		treeNode.Users[user.Id] = &user
	} else {
		for idx := range treeNode.Subordinates {
			subordinate := treeNode.Subordinates[idx]
			subordinate.InsertUser(user)
		}
	}
}

//Gets all subordinates, including subordinates of subordinates of user with given userID
func (tree *UserHierachyTree) GetSubordinates(userID int) []User {
	users := []User{}
	userRoleNode := tree.Root.FindTreeNodeByUserID(userID)
	if userRoleNode == nil {
		return users
	}
	if len(userRoleNode.Subordinates) > 0 {
		userRoleNode.FindSubordinateUsers(&users)
	}
	return users
}

// returns tree node if user is found in current node Users map
// if not found and there are subordinates
// recursively find user id in subordinates map
func (treeNode *TreeNode) FindTreeNodeByUserID(userID int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if treeNode.Users[userID] != nil {
		return treeNode
	} else if len(treeNode.Subordinates) > 0 {
		for idx := range treeNode.Subordinates {
			subordinate := treeNode.Subordinates[idx]
			return subordinate.FindTreeNodeByUserID(userID)
		}
	}
	return nil
}

// returns subordinates of current treeNode
// if treeNode is nil or no subordinates, return
// otherwise loop through subordinates and append users to list
func (treeNode *TreeNode) FindSubordinateUsers(users *[]User) {
	if treeNode == nil || len(treeNode.Subordinates) == 0 {
		return
	}
	for idx := range treeNode.Subordinates {
		subordinate := treeNode.Subordinates[idx]
		for _, user := range subordinate.Users {
			*users = append(*users, *user)
		}
		subordinate.FindSubordinateUsers(users)
	}
}

//sorts users by role in ascending order
func sortUsersByRole(users []User) {
	sort.Slice(users,
		func(i, j int) bool {
			return users[i].Role < users[j].Role
		})
}

//sorts roles by Parent in ascending order
func sortRolesByParent(roles []Role) {
	sort.Slice(roles,
		func(i, j int) bool {
			return roles[i].Parent < roles[j].Parent
		})
}
