package main

import (
	"ggymm/gorm-gen"
	"ggymm/gorm-gen/examples/dal/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	// generate from struct in project
	g.ApplyBasic(model.Customer{})

	g.Execute()
}
