#Build Stage
FROM golang:1.27.0-alpine AS builder

LABEL maintainer="mohamedamino882@gmail.com"  \
      version="0.1" \
      description="Backend API"


WORKDIR /build

COPY go.mod go.sum  ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/src ./cmd/server

#Runtime Stage
FROM alpine:3.21

LABEL maintainer="mohamedamino882@gmail.com"  \
      version="0.1" \
      description="Backend API"

WORKDIR /opt/app

RUN addgroup -S nonrootusers && adduser -S frappe -G nonrootusers

COPY --chown=frappe:nonrootusers --from=builder /build/src .

USER frappe:nonrootusers

ENTRYPOINT [ "./src" ]
