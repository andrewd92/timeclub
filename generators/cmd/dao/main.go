package main

import (
	"flag"
	"github.com/andrewd92/timeclub/generators/dao"
	log "github.com/sirupsen/logrus"
)

func main() {
	model := flag.String("m", "", "model path")
	table := flag.String("t", "", "table name")
	service := flag.String("s", "", "service folder")

	flag.Parse()

	if *model == "" {
		log.Fatal("empty model path")
	}
	if *table == "" {
		log.Fatal("empty table name")
	}
	if *service == "" {
		log.Fatal("empty service name")
	}

	dao.Generate(*model, *table, *service)
}
