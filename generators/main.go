package main

import (
	"flag"
	"github.com/andrewd92/timeclub/generators/dao"
	log "github.com/sirupsen/logrus"
)

func main() {
	table := flag.String("t", "", "table name")

	if *table == "" {
		log.Fatal("empty table name")
	}

	dao.Generate("model/model.go", *table, "club_service")
}
