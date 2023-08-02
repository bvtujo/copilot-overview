FROM golang:1.20 as builder

ENV GOPROXY=direct
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend ./pkg/frontend/main.go


FROM alpine:3.18 as run
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/frontend ./
EXPOSE 8080
CMD ["./frontend"]