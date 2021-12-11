package dao

import (
	"fmt"
	"github.com/gobeam/stringy"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"text/template"
)

type Data struct {
	Service      string
	Table        string
	Name         string
	Package      string
	Model        string
	Fields       string
	Placeholders string
	UpdateFields string
	FieldsMap    string
}

func Generate(path string, table string, service string) {
	columns := parse(path)

	placeholders := make([]string, len(columns), len(columns))
	set := make([]string, len(columns), len(columns))

	for i, column := range columns {
		placeholders[i] = ":" + column
		set[i] = column + " = :" + column
	}

	str := stringy.New(table)
	name := str.CamelCase()

	d := Data{
		Service:      service,
		Table:        fmt.Sprintf("`%s`", table),
		Name:         name,
		Package:      table + "_dao",
		Model:        fmt.Sprintf("model.%sModel", name),
		Fields:       strings.Join(columns, ", "),
		Placeholders: strings.Join(placeholders, ", "),
		UpdateFields: strings.Join(set, ", "),
	}

	modelPathParts := strings.Split(path, "/")
	infraPathParts := modelPathParts[:len(modelPathParts)-2]
	daoPathParts := append(infraPathParts, "dao", d.Package, "d.go")

	daoPath := strings.Join(daoPathParts, "/")

	f, err := os.Create(daoPath)
	if err != nil {
		log.WithError(err).Fatal("can not create file")
	}
	defer f.Close()

	t := template.Must(template.New("dao").Parse(templateStr))
	err = t.Execute(f, d)

	if err != nil {
		log.WithError(err).Error()
	}
}
