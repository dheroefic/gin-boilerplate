FROM golang:1.17-alpine

EXPOSE 1801

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod tidy && go mod vendor && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app

CMD ["app"]