package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type UserHierachyService interface {
	SetUsers(data []byte) []User
	SetRoles(data []byte) []Role
	GetSubordinates(userID int) []User
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

// store hierachy of users
type UserHierachyTree struct {
	Root *TreeNode
}

type IUserHierachyTree interface {
	FindTreeNodeByUserID(userID int) *TreeNode
}

type ITreeNode interface {
	InsertRole(role Role)
	InsertUser(user User)
	FindByUserID(userID int) *TreeNode
}

// each node stores
// Role, Users and Subordinates
// Users store users using map, with user id as key and user as Value
// allows more efficient check whether user is in current node's Role
// Subordinates stores list of TreeNode, assuming multiple roles can have the same parent
// Subordinate's role parent is current TreeNode
type TreeNode struct {
	Role         Role
	Users        map[int]*User
	Subordinates []*TreeNode
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
func (tree *UserHierachyTree) SetRoles(data []byte) {
	var roles []Role
	err := json.Unmarshal(data, &roles)
	if err != nil {
		fmt.Println("fail to unmarshal roles")
	}
	sortRolesByParent(roles)
	for idx := range roles {
		role := roles[idx]
		if tree.Root == nil {
			tree.Root = NewTreeNode(role)
		} else {
			tree.Root.InsertRole(role)
		}
	}
}

func (treeNode *TreeNode) InsertRole(role Role) {
	if treeNode.Role.Id == role.Parent {
		newNode := NewTreeNode(role)
		treeNode.Subordinates = append(treeNode.Subordinates, newNode)
	} else {
		for idx := range treeNode.Subordinates {
			subordinate := treeNode.Subordinates[idx]
			subordinate.InsertRole(role)
		}
	}
}

//sort users by role helpes constructing `UserHierachyTree`
func (tree *UserHierachyTree) SetUsers(data []byte) *UserHierachyTree {
	var users []User
	err := json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("fail to unmarshal roles")
	}
	sortUsersByRole(users)
	for idx := range users {
		user := users[idx]
		tree.Root.InsertUser(user)
	}
	return tree
}

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

func (tree *UserHierachyTree) GetSubordinates(userID int) []User {
	var users []User
	userRoleNode := tree.Root.FindTreeNodeByUserID(userID)
	if userRoleNode == nil {
		return users
	}
	if len(userRoleNode.Subordinates) > 0 {
		userRoleNode.FindSubordinates(&users)
	}
	return users
}

func (treeNode *TreeNode) FindTreeNodeByUserID(userID int) *TreeNode {
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

func (treeNode *TreeNode) FindSubordinates(users *[]User) {
	if treeNode == nil || len(treeNode.Users) == 0 {
		return
	}

	if len(treeNode.Subordinates) > 0 {
		for idx := range treeNode.Subordinates {
			subordinate := treeNode.Subordinates[idx]
			for _, user := range subordinate.Users {
				*users = append(*users, *user)
			}
			subordinate.FindSubordinates(users)
		}
	}
}

func sortUsersByRole(users []User) {
	sort.Slice(users,
		func(i, j int) bool {
			return users[i].Role < users[j].Role
		})
}

func sortRolesByParent(roles []Role) {
	sort.Slice(roles,
		func(i, j int) bool {
			return roles[i].Parent < roles[j].Parent
		})
}
