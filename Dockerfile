FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY ./ ./

#RUN go get
RUN go build -o main  .

CMD ["./main"]

