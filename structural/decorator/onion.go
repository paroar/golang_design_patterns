package decorator

import (
	"errors"
	"fmt"
)

type Onion struct {
	Ingredient IIngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil{
		return "", errors.New("Ingredient field on onion is empty")
	}
	pizza, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", pizza, "onion"), nil
}
