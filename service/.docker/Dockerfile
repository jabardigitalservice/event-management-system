# build
FROM --platform=linux/amd64 harbor.digitalservice.id/proxyjds/library/golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/services ./src/services
RUN CGO_ENABLED=0 go build -o bin/migration ./src/migration

FROM --platform=linux/amd64 harbor.digitalservice.id/proxyjds/library/alpine:3.16.0
COPY --from=builder /app/bin/services /
COPY --from=builder /app/bin/migration /
COPY --from=builder /app/.env.example /.env
COPY --from=builder /app/src /app
EXPOSE 30008
CMD ["/services"]