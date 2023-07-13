FROM golang:1.20.5

WORKDIR /app

COPY . /app

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]

