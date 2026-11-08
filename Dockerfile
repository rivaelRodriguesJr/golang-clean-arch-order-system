FROM golang:latest as builder
WORKDIR /app
COPY . .
WORKDIR /app/cmd/ordersystem

RUN CGO_ENABLED=0 go build -o server main.go wire_gen.go  

FROM scratch
COPY --from=builder /app/cmd/ordersystem/server .
CMD [ "./server" ]