package main

import (
	"beercli"
	"flag"
	"os"
)

func main() {

	all := flag.Bool("all", false, "View all available beers")
	random := flag.Bool("random", false, "View a randomly selected beer")
	byid := flag.Int("id", -1, "View beer with specified ID")

	flag.Parse()

	if *all != false {
		beer, err := beercli.GetAllBeers()
		if err != nil {
			panic(err)
		}

		for _, b := range beer {
			b.Pretty()
		}
		os.Exit(0)
	}

	if *random != false {
		randomBeer, err := beercli.GetRandomBeer()
		if err != nil {
			panic(err)
		}
		randomBeer.Pretty()
		os.Exit(0)
	}

	if *byid != -1 {
		beerByID, err := beercli.GetBeerByID(*byid)
		if err != nil {
			panic(err)
		}
		beerByID.Pretty()
		os.Exit(0)

	}

}
