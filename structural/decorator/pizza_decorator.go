package decorator

type PizzaDecorator struct {
	Ingredient IIngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with ingredients:", nil
}
