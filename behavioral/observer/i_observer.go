package observer

type IObserver interface {
	Notify(string)
}
