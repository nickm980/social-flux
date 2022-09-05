FROM golang:1.19-alpine3.15
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 10000
CMD [ "/app/main" ]