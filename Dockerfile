FROM golang:1.22.3

WORKDIR /app

COPY . .

RUN go build main.go

ENTRYPOINT ["/main"]