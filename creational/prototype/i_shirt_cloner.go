package prototype

type IShirtCloner interface{
	GetClone(s int) (ItemInfoGetter, error)
}