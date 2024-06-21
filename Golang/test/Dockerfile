FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

CMD ["./main"]

EXPOSE 3000

ENV PORT=3000
ENV JWT_SECRET=ssecrett
ENV DB_SOURCE=postgres://postgres:postgres@localhost/test?sslmode=disable
