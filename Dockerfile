FROM golang:latest AS builder
WORKDIR /app

COPY . .

RUN make install
RUN make build

FROM golang:latest
WORKDIR /app

COPY --from=builder /app/output/core /app/core
COPY --from=builder /app/config/ /app/config/

RUN chmod 755 /app/core

ENV GIN_MODE=release
ENTRYPOINT [ "./core" ]
