# Stage 1: build the program, with the full Go toolchain.
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /todoapi .

# Stage 2: only the finished program goes in the final image.
FROM alpine:3.22
WORKDIR /app
COPY --from=build /todoapi .
ENV DB_PATH=/data/todos.db
VOLUME /data
EXPOSE 8080
CMD ["./todoapi"]
