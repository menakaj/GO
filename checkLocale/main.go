package main

import (
	"fmt"
	"os"
)

type loacale struct {
	loc    string
	region string
}

func newLocale(loc string, region string) *loacale {
	l := loacale{loc: loc, region: region}
	return &l
}

func populateLocales() []loacale {
	l := []string{"en_FR", "en_EU", "en_US"}
	c := []string{"France", "UK", "USA"}
	var supportedLocales []loacale

	for i := 0; i < len(l); i++ {
		lc := newLocale(l[i], c[i])
		supportedLocales = append(supportedLocales, *lc)
	}
	return supportedLocales
}

func main() {

	lco := populateLocales()
	fmt.Println(lco)

	args := os.Args
	if len(args) < 3 {
		os.Exit(1)
	}

	for _, v := range lco {
		if v.loc == args[1] && v.region == args[2] {
			fmt.Println("locale supported")
		}
	}
}
