FROM golang:alpine
WORKDIR /usr/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main ./cmd
EXPOSE 80
CMD ["./main"]