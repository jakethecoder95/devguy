FROM golang:1.22 as builder
WORKDIR '/app'
COPY ./main.go ./
COPY ./index.html ./
RUN go build main.go

CMD ["./main"]
