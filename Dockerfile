ARG GOVERSION=1.16
FROM golang:${GOVERSION} as build-env

ARG GOARCH
ENV GOARCH=${GOARCH}

WORKDIR /go/src/voter-app
COPY . /go/src/voter-app

RUN go mod download
RUN go mod verify
RUN make build-local

# Copy into base image
FROM alpine:3.7
COPY --from=build-env /go/bin/voter-app /
COPY --from=build-env /go/src/voter-app/pkg /pkg
COPY --from=build-env /go/src/voter-app/views /views

RUN adduser -s /bin/false -h /home/voter -D voter

EXPOSE 8080

USER voter

CMD ["/voter-app"]
