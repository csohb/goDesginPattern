package obj

import "fmt"

type TeamMember struct {
	MemberName string
}

func (s *TeamMember) Listing() {
	fmt.Printf("This is TeamMember :%s\n", s.MemberName)
}

func (s *TeamMember) GetTeamMemberName() string {
	return s.MemberName
}
