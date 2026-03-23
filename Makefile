dev:
	npx nodemon --signal SIGTERM --exec go run ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

db:
	docker run -d --name fiberstarter-postgres \
		-e POSTGRES_USER=admin \
		-e POSTGRES_PASSWORD=admin \
		-p 5432:5432 \
		-v fiberstarter-postgres:/var/lib/postgresql \
		postgres


