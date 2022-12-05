FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /middleware-server

EXPOSE 8080

CMD ["/middleware-server"]
