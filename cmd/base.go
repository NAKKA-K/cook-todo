package main

import (
	"flag"

	"github.com/VG-Tech-Dojo/treasure2018/mid/NAKKA-K/cook-do"
)

func main() {
	var (
		addr   = flag.String("addr", ":8000", "addr to bind")
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)
	flag.Parse()
	b := base.New()
	b.Init(*dbconf, *env)
	b.Run(*addr)
}
