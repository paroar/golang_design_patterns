package decorator

import "testing"

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizzaDecorator := &PizzaDecorator{}
	ingredients, err := pizzaDecorator.AddIngredient()
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedMsg := "Pizza with ingredients:"
	if ingredients != expectedMsg {
		t.Errorf("Expected: %s\nActual: %s", expectedMsg, ingredients)
	}

}

func TestOnion_AddIngredient(t *testing.T) {

	t.Run("AddIngredient - Onion No Ingredient field", func(t *testing.T) {
		onion := &Onion{}
		_, err := onion.AddIngredient()
		if err == nil {
			t.Fatal("Calling AddIngredient on onion decorator without Ingredient field must return an error")
		}
	})

	t.Run("AddIngredient - Onion with Ingredient field", func(t *testing.T) {
		onion := &Onion{
			Ingredient: &PizzaDecorator{},
		}
		expectedMsg := "Pizza with ingredients: onion"
		ingredient, err := onion.AddIngredient()
		if err != nil {
			t.Fatal(err.Error())
		}
		if ingredient != expectedMsg {
			t.Errorf("Expected: %s\nActual: %s\n", expectedMsg, ingredient)
		}
	})
}

func TestMeatAddIngredient(t *testing.T) {
	t.Run("AddIngredient - Meat No Ingredient field", func(t *testing.T) {
		meat := &Meat{}
		_, err := meat.AddIngredient()
		if err == nil {
			t.Fatal("Calling AddIngredient on meat decorator without Ingredient field must return an error")
		}
	})

	t.Run("AddIngredient - Meat with Ingredient field", func(t *testing.T) {
		meat := &Meat{
			Ingredient: &PizzaDecorator{},
		}
		expectedMsg := "Pizza with ingredients: meat"
		ingredient, err := meat.AddIngredient()
		if err != nil {
			t.Fatal(err.Error())
		}
		if ingredient != expectedMsg {
			t.Errorf("Expected: %s\nActual: %s\n", expectedMsg, ingredient)
		}
	})
}

func TestOnionMeatAddIngredient(t *testing.T) {
	pizza := &Onion{
		Ingredient: &Meat{
			Ingredient: &PizzaDecorator{},
		},
	}
	expectedMsg := "Pizza with ingredients: meat onion"
	ingredient, err := pizza.AddIngredient()
	if err != nil {
		t.Fatal(err.Error())
	}
	if ingredient != expectedMsg {
		t.Errorf("Expected: %s\nActual: %s\n", expectedMsg, ingredient)
	}
}
