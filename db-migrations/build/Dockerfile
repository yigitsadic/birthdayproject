FROM golang:1.19-alpine

RUN apk add --no-cache make

WORKDIR /src

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# CMD migrate -path db -database "${POSTGRES_DSL}" up

RUN cd .. && touch .env && cd src

ENTRYPOINT "make"
CMD ["migrate"]

