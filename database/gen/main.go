package main

import (
	"gorm.io/gen"

	"github.com/star-horizon/anonymous-box-saas/database/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "database/dal",
		Mode:    gen.WithQueryInterface,
	})

	g.ApplyBasic(
		model.User{},
		model.Setting{},
		model.Website{},
		model.Comment{},
		model.Attachment{},
		model.Storage{},
	)

	g.Execute()
}
