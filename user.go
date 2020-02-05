package main

type Role struct {
	Id     int
	Name   string
	Parent int
}

type RoleNode struct {
	Value  Role
	Data   []User
	Parent *RoleNode
	Left   *RoleNode
	Right  *RoleNode
}

func BFS(roleTree *RoleNode, role int) []User {
	queue := []*RoleNode{}
	queue = append(queue, roleTree)
	result := []User{}
	return BFSUtil(queue, result, role)
	return []User{}
}

var root = 0

func BFSUtil(queue []*RoleNode, result []User, role int) []User {
	if len(queue) == 0 {
		return result
	}
	if queue[0].Left != nil && queue[0].Left.Value.Parent < role {

	}
	return BFSUtil(queue[1:], result)
}

type User struct {
	Id   int
	Name string
	Role int
}

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

// store roles in map, key -> ID, value -> role
// when setting users, create tree, rolesmap[user.role] -> gives the parent
// finally, whe get subordinates, find the user, return children
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
