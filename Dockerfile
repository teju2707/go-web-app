FROM golang:1.20-alpine AS base
WORKDIR /app
COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base
COPY --from=base /app/main .
COPY --from=base /app/static ./static
CMD ["./main"]