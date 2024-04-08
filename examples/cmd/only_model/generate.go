package main

import (
	"ggymm/gorm-gen"
	"ggymm/gorm-gen/examples/conf"
	"ggymm/gorm-gen/examples/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/tmp/gentest/query",
	})

	g.UseDB(dal.DB)

	g.GenerateAllTable()

	g.Execute()
}
