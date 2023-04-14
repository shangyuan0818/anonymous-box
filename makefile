gorm-gen:
	@echo "Generating gorm dal..."
	@go run database/gen/main.go

kitex-gen:
	@echo "Generating kitex idl..."

	@kitex -module github.com/star-horizon/anonymous-box-saas idl/base/timestamp.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/base/empty.proto

	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/auth.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/email.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/verify.proto

gen:
	@echo "Generating..."
	@make gorm-gen
	@make kitex-gen
	@echo "Done."
