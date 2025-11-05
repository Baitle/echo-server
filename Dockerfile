FROM golang:1.24 AS builder
ADD /app /app
WORKDIR /app
RUN go mod download
RUN go build -o /go/bin/main .

FROM gcr.io/distroless/base-debian12 AS production
WORKDIR /app
COPY --from=builder /go/bin/main /app/main
EXPOSE 3000
ENTRYPOINT ["/app/main"]