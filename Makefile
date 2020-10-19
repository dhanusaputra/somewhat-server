help: 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

proto: ## generate proto
	./third_party/protoc-gen.sh

mock: ## generate mocks from mockery
	mockery --all --inpkg

heroku: ## push to heroku
	git push heroku

keypair: ## generate key pair
	cd cert && openssl genrsa -out server.key 2048 && openssl req -new -x509 -key server.key -out server.pem -days 3650

.DEFAULT_GOAL := help