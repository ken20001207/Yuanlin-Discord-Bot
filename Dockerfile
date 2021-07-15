FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .
EXPOSE 9087
CMD ["/dist/main"]