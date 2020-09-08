package flyweight

import "testing"

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TeamA)
	if teamA1 == nil {
		t.Error("Pointer to TeamA was nil")
	}
	teamA2 := factory.GetTeam(TeamA)
	if teamA2 == nil {
		t.Error("Pointer to TeamA was nil")
	}
	if teamA1 != teamA2 {
		t.Error("Pointers to TeamA were different")
	}
	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()

	totalLength := 1000000
	teams := make([]*Team, totalLength)
	for i := 0; i < totalLength/2; i++ {
		teams[i] = factory.GetTeam(TeamA)
	}
	for i := totalLength / 2; i < totalLength; i++ {
		teams[i] = factory.GetTeam(TeamB)
	}
	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("Expected length to be 2 not %d", factory.GetNumberOfObjects())
	}
}
