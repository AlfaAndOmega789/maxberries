FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV PATH=$PATH:/go/bin

WORKDIR /app/catalog_service/cmd

RUN go build -o main .

CMD ["./main"]