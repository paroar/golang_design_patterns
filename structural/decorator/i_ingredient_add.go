package decorator

type IIngredientAdd interface {
	AddIngredient() (string, error)
}
