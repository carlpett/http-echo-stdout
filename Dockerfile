FROM golang:1.14 AS builder
COPY . .
RUN CGO_ENABLED=0 go build -o /server

FROM busybox
COPY --from=builder /server /server
CMD /server
