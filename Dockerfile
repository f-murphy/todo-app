FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./
COPY ./configs /configs

RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]