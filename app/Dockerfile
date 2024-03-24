FROM golang:1.22-alpine as builder
WORKDIR '/app'
COPY ./main.go ./
COPY ./index.html ./
RUN go build -o /bin/main main.go

FROM alpine:latest
COPY --from=builder /bin/main /bin/main
COPY --from=builder /app/index.html /index.html
CMD ["/bin/main"]

