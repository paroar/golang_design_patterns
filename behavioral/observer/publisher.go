package observer

type Publisher struct {
	ObserverList []IObserver
}

func (p *Publisher) AddObserver(o IObserver) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) RemoveObserver(o IObserver) {
	var indexToRemove int

	for i, observer := range p.ObserverList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	p.ObserverList = append(p.ObserverList[:indexToRemove], p.ObserverList[indexToRemove+1:]...)
}

func (p *Publisher) NotifyObservers(s string) {
	for _, o := range p.ObserverList {
		o.Notify(s)
	}
}
