package graph

import "github.com/damoon/ape-store/products/graph/model"

var price1 = 11
var price2 = 22
var price3 = 33

var hats = []*model.Product{
	{
		Upc:   "top-1",
		Name:  "Trilby",
		Price: &price1,
	},
	{
		Upc:   "top-2",
		Name:  "Fedora",
		Price: &price2,
	},
	{
		Upc:   "top-3",
		Name:  "Boater",
		Price: &price3,
	},
}
