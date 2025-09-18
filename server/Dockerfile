FROM golang:1.25 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go mod tidy
RUN go install github.com/cosmtrek/air@v1.51.0
RUN go build -o server .

FROM golang:1.25
WORKDIR /app
COPY --from=base /go/bin/air /usr/local/bin/air
COPY --from=base /app /app
CMD ["air"]
