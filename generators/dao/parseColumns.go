package dao

import (
	"fmt"
	"github.com/fatih/structtag"
	log "github.com/sirupsen/logrus"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func parse(path string) []string {
	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)

	if err != nil {
		log.WithError(err).Fatal("Error")
	}

	var columns = make([]string, 0)

	for _, f := range node.Decls {
		genD, ok := f.(*ast.GenDecl)
		if !ok {
			fmt.Printf("SKIP %T is not *ast.GenDecl\n", f)
			continue
		}

		for _, spec := range genD.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := currType.Type.(*ast.StructType)
			if !ok {
				continue
			}

			for _, field := range structType.Fields.List {
				tags, err := structtag.Parse(strings.Trim(field.Tag.Value, "`"))
				if err != nil {
					log.WithError(err).WithField("f", field.Names[0].Name).Fatal("Tag Err")
				}

				tag, err := tags.Get("db")

				if err != nil {
					log.WithError(err).Fatal("Db Tag Err")
				}

				if tag.Value() == "id" {
					continue
				}

				columns = append(columns, tag.Value())
			}
		}
	}

	return columns
}
