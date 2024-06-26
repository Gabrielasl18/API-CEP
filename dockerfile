FROM golang:1.22.3

WORKDIR /go-api-cep

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-api-cep

EXPOSE 8888

RUN "ls"

CMD ["./go-api-cep"]