package observer

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      int
	Message string
}

func (t *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received\n", t.ID, m)
	t.Message = m
}

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{ID: 1}
	testObserver2 := &TestObserver{ID: 2}
	testObserver3 := &TestObserver{ID: 3}
	publisher := &Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		if len(publisher.ObserverList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver1)

		if len(publisher.ObserverList) != 2 {
			t.Errorf("Size of observer list is not correct, expected 2, got %d", len(publisher.ObserverList))
		}

		for _, o := range publisher.ObserverList {
			observer, ok := o.(*TestObserver)
			if !ok {
				t.Fail()
			}
			if observer.ID == 1 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {
		if len(publisher.ObserverList) == 0 {
			t.Errorf("The list is empty. Nothing to test\n")
		}
		for _, o := range publisher.ObserverList {
			observer, ok := o.(*TestObserver)
			if !ok {
				t.Fail()
			}
			expected := ""
			if observer.Message != expected {
				t.Errorf("Expected: %s\nActual: %s", expected, observer.Message)
			}
		}

		publisher.NotifyObservers("Notifiying")
		for _, o := range publisher.ObserverList {
			observer, ok := o.(*TestObserver)
			if !ok {
				t.Fail()
			}
			expected := "Notifiying"
			if observer.Message != expected {
				t.Errorf("Expected: %s\nActual: %s", expected, observer.Message)
			}
		}
	})

}
