package main

import "goDesginPattern/composite/obj"

func main() {
	member1 := &obj.TeamMember{MemberName: "allen"}
	member2 := &obj.TeamMember{MemberName: "tony"}
	member3 := &obj.TeamMember{MemberName: "kevin"}
	member4 := &obj.TeamMember{MemberName: "dash"}
	member5 := &obj.TeamMember{MemberName: "winter"}
	member6 := &obj.TeamMember{MemberName: "noel"}

	devOps := &obj.Team{
		TeamName: "devOps",
	}

	devOps.AddTeamMember(*member1, *member2, *member4)

	tcpTeam := &obj.Team{
		TeamName: "tcp",
	}

	tcpTeam.AddTeamMember(*member3, *member5, *member6)

	k2Team := &obj.Team{
		TeamName: "K2",
	}

	k2Team.AddTeamMember(*member2)

	dev2Team := &obj.Team{
		TeamName: "개발2팀",
	}

	dev2Team.AddTeamMember(*member1, *member2, *member3, *member4, *member5, *member6)

	dev2Team.AddTeam(*devOps, *tcpTeam, *k2Team)

	dev2Team.Listing()

}
