FROM golang:alpine as builder
WORKDIR /app
RUN apk add ca-certificates tzdata
COPY . /app
RUN ls -all
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/app .

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/app /go/bin/app
ENV TZ=Asia/Jakarta
EXPOSE 8080
CMD ["/go/bin/app"]