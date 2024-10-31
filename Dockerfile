FROM golang:latest AS builder
WORKDIR /app

COPY . .

RUN make install
RUN make build

FROM golang:latest
WORKDIR /app

COPY --from=builder /app/output/receipt-api /app/receipt-api
COPY --from=builder /app/config/ /app/config/

RUN chmod 755 /app/receipt-api

ENV GIN_MODE=release
ENTRYPOINT [ "./core" ]
