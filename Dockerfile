ARG GO_VERSION

FROM golang:${GO_VERSION}
WORKDIR /logger
COPY . ./
RUN make install
RUN make test
