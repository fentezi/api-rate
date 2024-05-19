FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /rate ./cmd/


RUN wget -O migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz \
    && tar -xzf migrate.tar.gz -C /usr/local/bin migrate \
    && rm migrate.tar.gz


EXPOSE 8080

CMD ["sh", "-c", "migrate -path /api/internal/db/migrations -database $DATABASE_URL up && /rate"]
