FROM golang:1.21.1

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -trimpath -o /usr/local/bin/app

CMD ["app"]
