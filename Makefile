run:
	go run app/3.api/main.go

gen-swagger:
	cd app/3.api/ && swag init && cd ../..
