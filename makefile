gorm-gen:
	@echo "Generating gorm dal..."
	@go run database/gen/main.go

kitex-gen:
	@echo "Generating kitex idl..."

	@kitex -module github.com/star-horizon/anonymous-box-saas idl/base/timestamp.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/base/empty.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/base/pagination.proto

	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/dash/auth.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/dash/email.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/dash/verify.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/dash/website.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/dash/comment.proto

	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/box/website.proto
	@kitex -module github.com/star-horizon/anonymous-box-saas idl/api/box/comment.proto

gen:
	@echo "Generating..."
	@make gorm-gen
	@make kitex-gen
	@echo "Done."
