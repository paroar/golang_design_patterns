package prototype

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU: "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU: "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU: "empty",
	Color: Blue,
}
