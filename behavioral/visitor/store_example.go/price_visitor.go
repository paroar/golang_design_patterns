package main

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p IProductRetriever) {
	pv.Sum += p.GetPrice()
}
