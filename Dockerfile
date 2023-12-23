FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY template/ ./template/

RUN CGO_ENABLED=0 GOOS=linux go build -o /orchestrator-auth

ENV PORT=8080

CMD ["/orchestrator-auth"]