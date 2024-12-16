FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build main.go

RUN ls

ENTRYPOINT ["./main"]