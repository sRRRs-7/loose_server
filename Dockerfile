# Build stage
FROM golang:1.20-rc-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go get github.com/rubenv/sql-migrate
RUN go build -o main main.go

#Run stage
FROM golang:1.20-rc-alpine3.17
WORKDIR /app
COPY --from=builder /app .
COPY app.env .

RUN ["chmod", "+x", "start.sh"]
ENTRYPOINT ["start.sh"]

EXPOSE 8080
CMD ["sql-migrate up -config=sql_migrate.yml", "/app/main" ]