package test_glide_hierarchy

import (
	"github.com/sumory/idgen"
	"github.com/fatih/color"
)

func NextId() (error, int64) {
	color.Blue("Prints %s in blue.", "text")

	_, idWorker := idgen.NewIdWorker(1)
	return idWorker.NextId()
}