package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonStruct struct {
	image,
	name,
	price string
}

func main() {
	var pokemons []PokemonStruct;

	c := colly.NewCollector() 

	c.OnHTML("li.product", func (e *colly.HTMLElement) {
		pokemon := PokemonStruct{}

		pokemon.name = e.ChildText("h2");
		pokemon.image = e.ChildAttr("img", "src")
		pokemon.price = e.ChildText(".price")

		pokemons = append(pokemons, pokemon)
	})

	c.Visit("https://scrapeme.live/shop/")

	file, err := os.Create("pokemons.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() 

	writer := csv.NewWriter(file)

	headers := []string{
		"name",
		"image",
		"price",
	}

	writer.Write(headers)

	for _, pokemon := range pokemons {
		record := []string{
			pokemon.name,
			pokemon.image,
			pokemon.price,
		}

		writer.Write(record)
	}

	defer writer.Flush()
}