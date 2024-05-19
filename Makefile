run:
	docker-compose build --no-cache

up:
	migrate -path=internal/db/migrations -database "postgresql://root:secret@127.0.0.1:5438/rate?sslmode=disable" -verbose up

down:
	migrate -path=internal/db/migrations -database "postgresql://root:secret@127.0.0.1:5438/rate?sslmode=disable" -verbose down	

create:
	migrate create -ext=sql -dir=internal/db/migrations -seq=$(NAME)

.PHONY: run up down create	