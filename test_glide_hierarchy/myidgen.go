package test_glide_hierarchy

import (
	"github.com/sumory/idgen"
)

func NextId() (error, int64) {
	_, idWorker := idgen.NewIdWorker(1)
	return idWorker.NextId()
}