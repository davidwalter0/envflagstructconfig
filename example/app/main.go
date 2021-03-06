package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/davidwalter0/go-cfg"
	"github.com/davidwalter0/go-flag"
)

type Key string
type Value float64

// X flag
var X = flag.String("FLAG", "STRING_VALUE", "FLAG USAGE...")

func main() {

	{
		var myapp myAPP

		var sti *cfg.StructInfo = &cfg.StructInfo{
			StructPtr: &myapp,
		}
		if err := sti.Parse(); err != nil {
			log.Fatalf("%v\n", err)
		}

		log.Printf("%v %T\n", myapp, myapp)
		jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
		log.Printf("\n%v\n", string(jsonText))
		/*
			flag.Usage()
			// Error can't call parse again
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("\n\n\n\n")
					fmt.Printf("***Error*** %v\n", err)
					fmt.Printf("\n\n\n\n")
					// sti.Parse()
				}
			}()
			sti.Parse()
		*/
	}

	os.Exit(0)
}
