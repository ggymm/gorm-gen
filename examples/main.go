package main

import (
	"context"
	"fmt"

	"ggymm/gorm-gen/examples/biz"
	"ggymm/gorm-gen/examples/conf"
	"ggymm/gorm-gen/examples/dal"
	"ggymm/gorm-gen/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
