package main

import (
	model2 "github.com/star-horizon/anonymous-box-saas/database/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/database/dal",
		Mode:    gen.WithQueryInterface,
	})

	g.ApplyBasic(model2.User{}, model2.Setting{})

	g.Execute()
}
