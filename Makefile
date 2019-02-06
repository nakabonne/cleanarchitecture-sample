
run:
	cd src/cleanarchitecture-sample; GOOS=linux GOARCH=amd64 go build -o cleanarchitecture-sample
	docker-compose up
