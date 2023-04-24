FROM golang:1.19-alpine as builder

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o birthdayapi ./cmd/api 

FROM alpine

COPY --from=builder /src /.

CMD ./birthdayapi

