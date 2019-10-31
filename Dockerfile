FROM golang:1.13.1-alpine as builder
WORKDIR /go/src/github.com/spacepal/Spacepal-AI
RUN apk update \
      && apk upgrade \
      && apk add --no-cache ca-certificates \
      && apk add --no-cache git \
      && apk add --no-cache make \
      && apk add --no-cache curl
COPY . ./
RUN make

FROM scratch

WORKDIR /app

EXPOSE 3131
ENV PORT="3131"

COPY --from=builder /go/src/github.com/spacepal/Spacepal-AI/dist/aiservice /app/aiservice
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app/aiservice"]
