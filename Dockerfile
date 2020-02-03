ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}
WORKDIR /app
COPY . ./
RUN make install
RUN make test
