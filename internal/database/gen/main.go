package main

import (
	"gorm.io/gen"

	"github.com/ahdark-services/anonymous-box-saas/internal/database/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/database/dal",
		Mode:    gen.WithQueryInterface,
	})

	g.ApplyBasic(model.User{}, model.Setting{})

	g.Execute()
}
