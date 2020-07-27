FROM golang

COPY src /go/src/soma

WORKDIR /go/src/soma

RUN go build -o main.o

ENTRYPOINT ./main.o