FROM golang:1.22 as builder
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY ./ ./
RUN make build


FROM alpine
COPY --from=builder /go/bin/ /go/bin/
COPY --from=builder /go/scripts/ /go/scripts/
COPY --from=builder /go/migrations/ /go/migrations/
COPY --from=builder /go/Makefile /go/Makefile

ENV PATH="/go/bin:${PATH}"
RUN apk add --no-cache make ca-certificates tzdata libc6-compat dos2unix
RUN dos2unix /go/scripts/entrypoint.sh
RUN chmod +x /go/scripts/entrypoint.sh

ENTRYPOINT [ "/go/scripts/entrypoint.sh" ]

EXPOSE 8081
