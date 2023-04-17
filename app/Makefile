export GOOS=linux

.PHONY: clean
clean:
	@rm -rf ./bin/*

.PHONY: build
build:
	@ls lambda | xargs -I {} go build -o bin/{} lambda/{}/main.go

.PHONY: init
prod_init:
	@cd ../terraform && terraform init

.PHONY: deploy
prod_deploy: clean build
	@rm -f ../terraform/function.zip
	@cd ../terraform && terraform apply

.PHONY: destroy
prod_destroy:
	@cd ../terraform && terraform destroy