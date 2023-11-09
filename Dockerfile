FROM golang:1.21.4 as base

FROM base as dev

RUN go install github.com/bokwoon95/wgo@latest

WORKDIR /app/api

CMD ["wgo", "run", "main.go"]

EXPOSE 8080