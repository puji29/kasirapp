FROM golang:alpine

WORKDIR /app

COPY . .
COPY .env /app

RUN go mod tidy
RUN go build -o kasir-app

ENTRYPOINT ["/app/kasir-app"]
