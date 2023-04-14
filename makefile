gorm-gen:
	@echo "Generating gorm dal..."
	@go run database/gen/main.go

kitex-gen:
	@echo "Generating kitex idl..."
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/auth.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/email.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/verify.proto
