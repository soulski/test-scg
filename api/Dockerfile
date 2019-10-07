FROM golang:1.13 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/app/build/scg .

ENTRYPOINT ["/app/scg"]
CMD ["start"]

EXPOSE 8080
