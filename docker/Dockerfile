# -- multistage docker build: stage #1: build stage
FROM golang:1.23-alpine AS build

RUN mkdir -p /go/src/github.com/bitmeme-taxi/bitmemed

WORKDIR /go/src/github.com/bitmeme-taxi/bitmemed

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache kaspad dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o kaspad .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/bitmeme-taxi/bitmemed/kaspad /app/
COPY --from=build /go/src/github.com/bitmeme-taxi/bitmemed/infrastructure/config/sample-kaspad.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
