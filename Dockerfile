# >-----< BASE IMAGE >-----< #

FROM golang:1.24-alpine AS base

# >-----< INSTALL STAGE >-----< #

FROM base AS installer

WORKDIR /app/

COPY \
  go.mod \
  go.sum ./

RUN go mod download

# >-----< BUILD STAGE >-----< #

FROM installer AS builder

ENV \
  CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /app/

COPY source/ source/

RUN go build -ldflags="-s -w" -o /app/termdo-gateway-api ./source/main.go

# >-----< RUN STAGE >-----< #

FROM scratch AS runner

WORKDIR /app/

COPY --from=builder /app/termdo-gateway-api termdo-gateway-api

ENTRYPOINT [ "./termdo-gateway-api" ]
