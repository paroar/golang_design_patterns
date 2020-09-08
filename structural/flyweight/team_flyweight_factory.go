package flyweight

type TeamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *TeamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *TeamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TeamB:
		return Team{
			ID:   2,
			Name: TeamB,
		}
	default:
		return Team{
			ID:   1,
			Name: TeamA,
		}
	}
}

func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}
