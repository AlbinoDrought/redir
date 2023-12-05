FROM golang:1.21 AS builder
WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod main.go /app/
RUN go get \
  && mkdir -p /app/dist \
  && go build -o /app/dist/redir \
  && echo "nobody:*:65534:65534:nobody:/_nonexistent:/bin/false" > /app/dist/passwd

FROM scratch
COPY --from=builder /app/dist/passwd /etc/passwd
COPY --from=builder /app/dist/redir /redir
USER nobody
ENTRYPOINT ["/redir"]
