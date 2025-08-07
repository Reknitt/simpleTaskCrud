FROM golang:alpine as build
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/taskCrud/.

FROM alpine:latest
WORKDIR /app

COPY --from=build /app/main /bin/main
ENTRYPOINT ["/bin/main"]
