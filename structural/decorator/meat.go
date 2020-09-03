package decorator

import (
	"errors"
	"fmt"
)

type Meat struct {
	Ingredient IIngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("Ingredient field on meat is empty")
	}
	pizza, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", pizza, "meat"), nil
}
