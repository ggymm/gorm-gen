package tests_test

import (
	"context"
	"sync"

	"ggymm/gorm-gen/tests/.expect/dal_test/query"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	query.Use(DB)
	query.SetDefault(DB)
}
