package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type PokemonStruct struct {
	url,
	image,
	name,
	price string
}

func main() {
	var pokemons []PokemonStruct;

	c := colly.NewCollector() 
 
	c.Visit("https://scrapeme.live/shop/")

	c.OnHTML("nav", func (e *colly.HTMLElement) {
		pokemon := PokemonStruct{}

		pokemon.name = e.ChildText("span")

		pokemons = append(pokemons, pokemon)
	})

	fmt.Printf("%v", pokemons)
}