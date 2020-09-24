package memento

import "testing"

func TestCareTaker_Add(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	careTaker := careTaker{}
	mem := originator.NewMemento()
	if mem.state.Description != "Idle" {
		t.Errorf("Expected: 'Idle', got %s", mem.state.Description)
	}

	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)

	if len(careTaker.mementoList) != currentLen+1 {
		t.Error("Expected mementoList increment by 1")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	careTaker := careTaker{}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Error(err)
	}

	if mem.state.Description != "Idle" {
		t.Errorf("Expected 'Idle', got %s", mem.state.Description)
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("Expected not found error")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	careTaker := careTaker{}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Error(err)
	}

	if mem.state.Description != "Idle" {
		t.Errorf("Expected 'Idle', got %s", mem.state.Description)
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("Expected not found error")
	}
}
