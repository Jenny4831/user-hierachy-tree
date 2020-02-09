package main

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// for true {
	// 	input, inputErr := reader.ReadString('\n')
	// 	if inputErr != nil {
	// 		fmt.Fprintln(os.Stderr, inputErr)
	// 	}

	// 	command := CoverStringToCommand(input)
	// userHierachy := NewUserHierachy()

	//set roles
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
	// var userHierachyTree *UserHierachyTree
	userHierachyTree := &UserHierachyTree{}
	userHierachyTree.SetRoles(rolesRequest)

}
