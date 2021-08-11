package graph

import "github.com/damoon/ape-store/reviews/graph/model"

var reviews = []*model.Review{
	{
		Body:    "A highly effective form of birth control.",
		Product: &model.Product{Upc: "top-1"},
		Author:  &model.User{ID: "1234", Username: "U1234"},
	},
	{
		Body:    "Fedoras are one of the most fashionable hats around and can look great with a variety of outfits.",
		Product: &model.Product{Upc: "top-1"},
		Author:  &model.User{ID: "1234", Username: "U1234"},
	},
	{
		Body:    "This is the last straw. Hat you will wear. 11/10",
		Product: &model.Product{Upc: "top-1"},
		Author:  &model.User{ID: "7777", Username: "U1234"},
	},
}
