package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"1": "Estrella Damm",
	"2": "Moritz",
	"3": "Heineken",
}

func main() {
	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("You must specify argument --beers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "find your beer")
		beersCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}

}
