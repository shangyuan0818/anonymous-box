package main

import (
	"gorm.io/gen"

	"github.com/star-horizon/anonymous-box-saas/database/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/database/dal",
		Mode:    gen.WithQueryInterface,
	})

	g.ApplyBasic(model.User{}, model.Setting{})

	g.Execute()
}
