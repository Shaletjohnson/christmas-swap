package giftassignment

var group *Group

//Group type will allow the user to add a group and set a budget
type Group struct {
	Name   string
	Budget float64

	persons         []string
	yearAssignments []YearAssignment
}

// Assignment New type to match gift givers and receivers
type Assignment struct {
	Giver    string
	Receiver string
}

// YearAssignment type to define what is needed to rotate through assignments
type YearAssignment struct {
	Year        int
	Assignments []Assignment
}

// AddPerson is to add people to your list for func Assign to handle
func (g *Group) AddPerson(person string) {
	g.persons = append(g.persons, person)
}

// Assign the givers and receivers using Assignment struct
// i year index j person index
func (g *Group) Assign() {
	for i := 0; i < len(g.persons)-1; i++ {
		ya := YearAssignment{
			Year: 2019 + i,
		}
		for j, giver := range g.persons {

			ri := (i + j + 1) % len(g.persons)
			receiver := g.persons[ri]
			assignment := Assignment{
				Giver:    giver,
				Receiver: receiver,
			}

			//put giver+receiver into assignment struct, and then add assignment to year assignment
			ya.Assignments = append(ya.Assignments, assignment)

		}
		//need to append to the yearAssignment array to use as a variable
		g.yearAssignments = append(g.yearAssignments, ya)
	}

}

//ListAssignments will allow us to list the array when called in the CLI
func (g *Group) ListAssignments() []YearAssignment {
	return g.yearAssignments
}

//SetGroup allows CLI to ask user to set up a group
func SetGroup(g *Group) {
	group = g
}

//GetGroup allows CLI to get the group when called so it can be printed for the user
func GetGroup() *Group {
	return group
}
