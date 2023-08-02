FROM golang:1.20 as builder

ENV GOPROXY=direct
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend .


FROM alpine:3.18 as run
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app ./
EXPOSE 8080
CMD ["./frontend"]