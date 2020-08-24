package prototype

import "testing"

func TestShirtCloner(t *testing.T) {
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatalf("Received cache was nil")
	}
}

func TestShirtNonExistent(t *testing.T) {
	shirtCloner := GetShirtCloner()
	_, err := shirtCloner.GetClone(-1)
	if err == nil {
		t.Error("Expected non existent Shirt error")
	}
}

func TestTypeAssertionShirt(t *testing.T){
	shirtCloner := GetShirtCloner()

	//protos & shirts must be in the same order
	var protos = []*Shirt{whitePrototype, blackPrototype, bluePrototype}
	var shirts = []int{White, Black, Blue}
	for i, v := range shirts {
		proto, err := shirtCloner.GetClone(v)
		if err != nil{
			t.Fatal(err)
		}
		shirt, ok := proto.(*Shirt)
		if !ok {
			t.Fatalf("Type assertion couldn't be done for Shirt id %d", v)
		}
		original := protos[i]
		if shirt == original {
			t.Errorf("Expected new Shirt to be different to the prototype %d %d", &shirt, &original)
		}
	}
}

func TestGetPrice(t *testing.T) {
	shirtCloner := GetShirtCloner()

	//protos & shirts must be in the same order
	var protos = []*Shirt{whitePrototype, blackPrototype, bluePrototype}
	var shirts = []int{White, Black, Blue}
	for i, v := range shirts {
		proto, _ := shirtCloner.GetClone(v)
		shirt := proto.(*Shirt)
		price := shirt.GetPrice()
		original := protos[i]
		if price != original.GetPrice() {
			t.Errorf("Expected %f, instead got %f", original.GetPrice(), price)
		}
	}
}