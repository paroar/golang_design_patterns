package singleton

import "testing"

func TestGetInstance(t *testing.T){
	counter1 := GetInstance()

	if counter1 == nil {
		t.Errorf("expected pointer to Singleton after calling GetInstance(), not nill\n")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("Expected the counter to be 1, intead it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Errorf("Expected same instance in counter2 but got a different instance\n")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("Expected current count 2 instead of %d\n", currentCount)
	}





}