package prototype

func GetShirtCloner() IShirtCloner{
	return new(ShirtCache)
}