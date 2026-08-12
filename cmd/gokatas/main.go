// Print statistics about katas you've done.
package main

import (
	"flag"
	"log"
	"os"
	"path"
)

func main() {
	log.SetPrefix(path.Base(os.Args[0]) + ": ")
	log.SetFlags(0)

	flag.Usage = func() { flag.PrintDefaults() }
	sortByColumn := flag.Int("c", 1, "sort katas by `column`")
	gokatasRepo := flag.String("r", ".", "path to gokatas repository")
	flag.Parse()

	if *gokatasRepo != "." {
		if err := os.Chdir(*gokatasRepo); err != nil {
			log.Fatal(err)
		}
	}

	katas, err := Get()
	if err != nil {
		log.Fatalf("getting katas: %v", err)
	}
	Print(katas, *sortByColumn)
}
