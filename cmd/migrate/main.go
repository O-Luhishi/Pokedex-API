package main

import (
	"flag"
	"fmt"
	"github.com/lunetco/pokedex_api/adapter"
	"github.com/lunetco/pokedex_api/config"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

const dialect = "postgres"

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "/app/migrations", "directory with migration files")
)

func main() {
	flags.Usage = usage
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing args: %v", err)
	}
	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}
	command := args[0]
	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	}
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	appDb, err := adapter.New(appConfig.Database)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer appDb.Close()
	if err := goose.SetDialect(dialect); err != nil {
		log.Fatal(err)
	}
	if err := goose.Run(command, appDb.DB(), *dir, args[1:]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}
func usage() {
	fmt.Print(usagePrefix)
	flags.PrintDefaults()
	fmt.Print(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
    migrate status
Options:
`
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
