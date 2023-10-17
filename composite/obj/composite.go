package obj

import "fmt"

type Team struct {
	Teams       []Team
	TeamMembers []TeamMember
	TeamName    string
}

func (t *Team) Listing() {
	fmt.Printf("%s has below team members\n", t.TeamName)
	fmt.Printf("teamMember:%s\n", t.TeamMembers)
	for _, composite := range t.Teams {
		composite.Listing()
		for _, member := range composite.TeamMembers {
			member.Listing()
		}
	}
}

func (t *Team) AddTeam(teams ...Team) {
	for _, team := range teams {
		t.Teams = append(t.Teams, team)
	}
}

func (t *Team) AddTeamMember(members ...TeamMember) {
	for _, member := range members {
		t.TeamMembers = append(t.TeamMembers, member)
	}
}
