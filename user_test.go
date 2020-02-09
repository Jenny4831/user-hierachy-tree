package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TreeRequest struct {
	Root              *TreeNode
	FirstSubordinate  *TreeNode
	SecondSubordinate *TreeNode
}

func defaultTreeForTest(treeRequest TreeRequest) *UserHierachyTree {
	userHierachyTree := &UserHierachyTree{}
	root := NewTreeNode(Role{
		Id:     1,
		Name:   "System Administrator",
		Parent: 0,
	})
	firstSubordinate := NewTreeNode(Role{
		Id:     2,
		Name:   "Location Manager",
		Parent: 1,
	})
	secondSubordinate := NewTreeNode(Role{
		Id:     3,
		Name:   "Supervisor",
		Parent: 2,
	})
	// if treeRequest.Root != nil {
	// 	root = treeRequest.Root
	// }
	// if treeRequest.FirstSubordinates !=nil {
	// 	firstSubordinates = treeRequest.FirstSubordinates
	// }
	// if treeRequest.SecondSubordinates != nil {
	// 	secondSubordinates = treeRequest.SecondSubordinates
	// }
	userHierachyTree.Root = root
	userHierachyTree.Root.Subordinates = []*TreeNode{firstSubordinate}
	firstSubordinate.Subordinates = []*TreeNode{secondSubordinate}
	return userHierachyTree
}

func TestUserHierachyTree_SetRoles(t *testing.T) {
	type fields struct {
		Root *TreeNode
	}

	tests := []struct {
		name    string
		tree    *UserHierachyTree
		roles   []Role
		want    *UserHierachyTree
		wantErr error
	}{
		{
			name: "sorted list of roles",
			tree: &UserHierachyTree{},
			roles: []Role{
				{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				{
					Id:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
				{
					Id:     3,
					Name:   "Supervisor",
					Parent: 2,
				},
			},
			want:    defaultTreeForTest(TreeRequest{}),
			wantErr: nil,
		},
		{
			name: "unsorted roles list",
			tree: &UserHierachyTree{},
			roles: []Role{
				{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				{
					Id:     3,
					Name:   "Supervisor",
					Parent: 2,
				},
				{
					Id:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
			},
			want:    defaultTreeForTest(TreeRequest{}),
			wantErr: nil,
		},
		{
			name: "multiple subordinates with same parent",
			tree: &UserHierachyTree{},
			roles: []Role{
				{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				{
					Id:     3,
					Name:   "Supervisor",
					Parent: 1,
				},
				{
					Id:     2,
					Name:   "Brand Manager",
					Parent: 1,
				},
			},
			want: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: make(map[int]*User),
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     3,
								Name:   "Supervisor",
								Parent: 1,
							},
							Users:        make(map[int]*User),
							Subordinates: []*TreeNode{},
						},
						{
							Role: Role{
								Id:     2,
								Name:   "Brand Manager",
								Parent: 1,
							},
							Users:        make(map[int]*User),
							Subordinates: []*TreeNode{},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:    "no roles in list",
			tree:    &UserHierachyTree{},
			roles:   []Role{},
			want:    &UserHierachyTree{},
			wantErr: fmt.Errorf("list of roles is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.tree.SetRoles(tt.roles)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("%s returned %+v, want %+v", tt.name, gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.tree, tt.want) {
				t.Errorf("%s returned %+v, want %+v", tt.name, tt.tree, tt.want)
			}
		})
	}
}

func TestUserHierachyTree_SetUsers(t *testing.T) {
	tests := []struct {
		name    string
		tree    *UserHierachyTree
		users   []User
		want    *UserHierachyTree
		wantErr error
	}{
		{
			name: "sorted list of users",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: make(map[int]*User),
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: make(map[int]*User),
							Subordinates: []*TreeNode{
								{Role: Role{
									Id:     3,
									Name:   "Supervisor",
									Parent: 2,
								},
									Users:        make(map[int]*User),
									Subordinates: []*TreeNode{}},
							},
						},
					},
				},
			},
			users: []User{
				{Id: 1, Name: "Adam Admin", Role: 1},
				{Id: 2, Name: "Sam Supervisor", Role: 3},
				{Id: 3, Name: "Mary Manager", Role: 2},
			},
			want: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: map[int]*User{
						1: &User{Id: 1, Name: "Adam Admin", Role: 1},
					},
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: map[int]*User{
								3: &User{Id: 3, Name: "Mary Manager", Role: 2},
							},
							Subordinates: []*TreeNode{
								{
									Role: Role{
										Id:     3,
										Name:   "Supervisor",
										Parent: 2,
									},
									Users: map[int]*User{
										2: &User{Id: 2, Name: "Sam Supervisor", Role: 3},
									},
									Subordinates: []*TreeNode{}},
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "role does not exist in tree",
			tree: &UserHierachyTree{Root: &TreeNode{
				Role: Role{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				Users: map[int]*User{
					1: &User{Id: 1, Name: "Adam Admin", Role: 1},
				},
				Subordinates: []*TreeNode{},
			}},
			users: []User{
				{Id: 2, Name: "Sam Supervisor", Role: 3},
			},
			want: &UserHierachyTree{Root: &TreeNode{
				Role: Role{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				Users: map[int]*User{
					1: &User{Id: 1, Name: "Adam Admin", Role: 1},
				},
				Subordinates: []*TreeNode{},
			}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.tree.SetUsers(tt.users)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("%s returned %+v, want %+v", tt.name, gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.tree, tt.want) {
				t.Errorf("%s returned %+v, want %+v", tt.name, tt.tree, tt.want)
			}
		})
	}
}

func TestUserHierachyTree_GetSubordinates(t *testing.T) {
	tests := []struct {
		name   string
		tree   *UserHierachyTree
		userID int
		want   []User
	}{
		{
			name: "Get subordinates, roles with different parents",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: map[int]*User{
						1: &User{Id: 1, Name: "Adam Admin", Role: 1},
					},
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: map[int]*User{
								3: &User{Id: 3, Name: "Mary Manager", Role: 2},
							},
							Subordinates: []*TreeNode{
								{
									Role: Role{
										Id:     3,
										Name:   "Supervisor",
										Parent: 2,
									},
									Users: map[int]*User{
										2: &User{Id: 2, Name: "Sam Supervisor", Role: 3},
									},
									Subordinates: []*TreeNode{}},
							},
						},
					},
				},
			},
			userID: 1,
			want: []User{
				{Id: 3, Name: "Mary Manager", Role: 2},
				{Id: 2, Name: "Sam Supervisor", Role: 3},
			},
		},
		{
			name: "Multiple subordinates, same parent, different roles",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: map[int]*User{
						1: &User{Id: 1, Name: "Adam Admin", Role: 1},
					},
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: map[int]*User{
								3: &User{Id: 3, Name: "Mary Manager", Role: 2},
							},
							Subordinates: []*TreeNode{},
						},
						{
							Role: Role{
								Id:     3,
								Name:   "Supervisor",
								Parent: 2,
							},
							Users: map[int]*User{
								2: &User{Id: 2, Name: "Sam Supervisor", Role: 3},
							},
							Subordinates: []*TreeNode{},
						},
					},
				},
			},
			userID: 1,
			want: []User{
				{Id: 3, Name: "Mary Manager", Role: 2},
				{Id: 2, Name: "Sam Supervisor", Role: 3},
			},
		},
		{
			name: "User in tree does not exist, empty list",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: map[int]*User{
						1: &User{Id: 1, Name: "Adam Admin", Role: 1},
					},
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: map[int]*User{
								3: &User{Id: 3, Name: "Mary Manager", Role: 2},
							},
							Subordinates: []*TreeNode{
								{
									Role: Role{
										Id:     3,
										Name:   "Supervisor",
										Parent: 2,
									},
									Users: map[int]*User{
										2: &User{Id: 2, Name: "Sam Supervisor", Role: 3},
									},
									Subordinates: []*TreeNode{}},
							},
						},
					},
				},
			},
			userID: 4,
			want:   []User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.GetSubordinates(tt.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserHierachyTree.GetSubordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_FindTreeNodeByUserID(t *testing.T) {
	tests := []struct {
		name     string
		rootNode *TreeNode
		userID   int
		want     *TreeNode
	}{
		{
			name: "Find tree node, user exists",
			rootNode: &TreeNode{
				Role: Role{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				Users: make(map[int]*User),
				Subordinates: []*TreeNode{{
					Role: Role{
						Id:     2,
						Name:   "Location Manager",
						Parent: 1,
					},
					Users: map[int]*User{
						3: &User{Id: 3, Name: "Mary Manager", Role: 2},
					},
				}},
			},
			userID: 3,
			want: &TreeNode{
				Role: Role{
					Id:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
				Users: map[int]*User{
					3: &User{
						Id: 3, Name: "Mary Manager", Role: 2,
					}},
			},
		},
		{
			name: "Find tree node by user ID, user does not exists",
			rootNode: &TreeNode{
				Role: Role{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				Users:        make(map[int]*User),
				Subordinates: []*TreeNode{}},
			userID: 2,
			want:   nil,
		},
		{
			name:     "Find tree node by user ID, user hierachy root is nil",
			rootNode: nil,
			userID:   2,
			want:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rootNode.FindTreeNodeByUserID(tt.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.FindTreeNodeByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortUsersByRole(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortUsersByRole(tt.args.users)
		})
	}
}

func Test_sortRolesByParent(t *testing.T) {
	tests := []struct {
		name  string
		roles []Role
		want  []Role
	}{
		{
			name: "ordered list",
			roles: []Role{
				{Id: 1, Name: "Admin", Parent: 0},
				{Id: 2, Name: "Supervisor", Parent: 1},
				{Id: 3, Name: "Manager", Parent: 2},
			},
			want: []Role{
				{Id: 1, Name: "Admin", Parent: 0},
				{Id: 2, Name: "Supervisor", Parent: 1},
				{Id: 3, Name: "Manager", Parent: 2},
			},
		},
		{
			name: "unordered list",
			roles: []Role{
				{Id: 3, Name: "Manager", Parent: 2},
				{Id: 1, Name: "Admin", Parent: 0},
				{Id: 2, Name: "Supervisor", Parent: 1},
			},
			want: []Role{
				{Id: 1, Name: "Admin", Parent: 0},
				{Id: 2, Name: "Supervisor", Parent: 1},
				{Id: 3, Name: "Manager", Parent: 2},
			},
		},
		{
			name: "unordered list, same parent",
			roles: []Role{
				{Id: 2, Name: "Supervisor", Parent: 1},
				{Id: 3, Name: "Manager", Parent: 1},
				{Id: 1, Name: "Admin", Parent: 0},
			},
			want: []Role{
				{Id: 1, Name: "Admin", Parent: 0},
				{Id: 2, Name: "Supervisor", Parent: 1},
				{Id: 3, Name: "Manager", Parent: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortRolesByParent(tt.roles)
		})
	}
}

func TestTreeNode_FindSubordinateUsers(t *testing.T) {
	tests := []struct {
		name string
		tree *UserHierachyTree
		want []User
	}{
		{
			name: "Get subordinates, one level hierachy",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: make(map[int]*User),
					Subordinates: []*TreeNode{{
						Role: Role{
							Id:     2,
							Name:   "Location Manager",
							Parent: 1,
						},
						Users: map[int]*User{
							3: &User{Id: 3, Name: "Mary Manager", Role: 2},
						},
						Subordinates: []*TreeNode{},
					}},
				},
			},
			want: []User{{Id: 3, Name: "Mary Manager", Role: 2}},
		},
		{
			name: "Get multiple subordinates, one level hierachy",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: map[int]*User{1: &User{Id: 1, Name: "Adam Admin", Role: 1}},
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: map[int]*User{
								3: &User{Id: 3, Name: "Mary Manager", Role: 2},
							},
							Subordinates: []*TreeNode{},
						},
						{
							Role: Role{
								Id:     3,
								Name:   "Supervisor",
								Parent: 1,
							},
							Users: map[int]*User{
								2: &User{Id: 2, Name: "Sam Supervisor", Role: 3},
							},
							Subordinates: []*TreeNode{},
						},
					},
				},
			},
			want: []User{
				{Id: 3, Name: "Mary Manager", Role: 2},
				{Id: 2, Name: "Sam Supervisor", Role: 3},
			},
		},
		{
			name: "Get multiple subordinates, multi lever hierachy",
			tree: &UserHierachyTree{
				Root: &TreeNode{
					Role: Role{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					Users: map[int]*User{
						1: &User{Id: 1, Name: "Mary Manager", Role: 2},
					},
					Subordinates: []*TreeNode{
						{
							Role: Role{
								Id:     2,
								Name:   "Location Manager",
								Parent: 1,
							},
							Users: map[int]*User{
								3: &User{Id: 3, Name: "Mary Manager", Role: 2},
							},
							Subordinates: []*TreeNode{
								{
									Role: Role{
										Id:     4,
										Name:   "Assistant",
										Parent: 2,
									},
									Users: map[int]*User{
										4: &User{Id: 4, Name: "Ella Assistant", Role: 4},
									},
									Subordinates: []*TreeNode{},
								},
							},
						},
						{
							Role: Role{
								Id:     3,
								Name:   "Supervisor",
								Parent: 1,
							},
							Users: map[int]*User{
								2: &User{Id: 2, Name: "Sam Supervisor", Role: 3},
							},
							Subordinates: []*TreeNode{},
						},
					},
				},
			},
			want: []User{
				{Id: 3, Name: "Mary Manager", Role: 2},
				{Id: 4, Name: "Ella Assistant", Role: 4},
				{Id: 2, Name: "Sam Supervisor", Role: 3},
			},
		},
		{
			name: "root is nil, empty list",
			tree: &UserHierachyTree{Root: nil},
			want: []User{},
		},
		{
			name: "no subordinates",
			tree: &UserHierachyTree{Root: &TreeNode{
				Role: Role{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				Users: map[int]*User{
					1: &User{Id: 1, Name: "Mary Manager", Role: 2},
				},
				Subordinates: []*TreeNode{},
			}},
			want: []User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []User{}
			tt.tree.Root.FindSubordinateUsers(&got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.FindTreeNodeByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
