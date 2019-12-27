ARG GO_VERSION

FROM golang:${GO_VERSION} as builder
WORKDIR /logger
COPY . ./
RUN make install
RUN make test
