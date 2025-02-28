FROM golang:1.23rc1-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.20 
WORKDIR /app
COPY --from=builder /app/main .
COPY ./data /app/data
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]