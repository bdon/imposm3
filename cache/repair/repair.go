package repair

import (
  "flag"
	"fmt"
	"log"
	"os"

  "github.com/jmhodges/levigo"
)

var flags = flag.NewFlagSet("repair-cache", flag.ExitOnError)
var cachedir = flags.String("cachedir", "/tmp/imposm3", "cache directory")

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s %s:\n\n", os.Args[0], os.Args[1])
	flags.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nRepair cache.")
	os.Exit(1)
}

func Repair(args []string) {
	flags.Usage = Usage

	if len(args) == 0 {
		Usage()
	}

	err := flags.Parse(args)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

  opts := levigo.NewOptions()
  fmt.Fprintf(os.Stderr, "Repairing %s", *cachedir)
  err = levigo.RepairDatabase(*cachedir,opts)
  if err != nil {
    log.Fatal(err)
  }
}
