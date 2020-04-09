package model

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
)

// From https://github.com/ekhabarov/go-pg-generator

func GenerateModel(options Options, tables ...string) {
	if options.User == "" {
		//obtain current logged user
		u, err := user.Current()
		if err != nil {
			log.Fatalln("unable to get current os user: ", err)
		}
		options.User = u.Username
	}

	db, err := dbConnect(options)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	defer db.Close()

	p := "models"
	if options.PackageName != "" {
		p = options.PackageName
	}
	pl := fmt.Sprintf("package %s \n", p)

	result := ""

	if len(tables) == 0 {
		tables, err = tableList(db)
		if err != nil {
			log.Fatalln(err)
		}
	}

	var mergedCols []*column
	for _, t := range tables {
		cols, err := columnList(db, t)
		if err != nil {
			log.Fatalln(err)
		}

		data := getStruct(t, cols)
		imports := getImports(cols)

		if options.FilePerTable {
			data = imports + data
		} else {
			mergedCols = append(mergedCols, cols...)
		}

		if options.FilePerTable {
			if err := saveToFile(filepath.Join(options.Dir, t), []byte(pl+data)); err != nil {
				log.Fatalln(err)
			}
			continue
		} else {
			result += data
		}
	}

	if !options.FilePerTable {
		imports := getImports(mergedCols)
		result = pl + imports + result
		if err := saveToFile(filepath.Join(options.Dir, options.OneFileName), []byte(result)); err != nil {
			log.Fatalln(err)
		}
	}
}
