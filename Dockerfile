# Divide into multiple stages to minimize the size of the image
# Build stage
FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go
# Install curl
RUN apk add curl
# Then download golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/main .
# Copy golang-migrate installing from builder stage to run stage
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
# Then copy the migration files 
COPY db/migration ./migration

EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]