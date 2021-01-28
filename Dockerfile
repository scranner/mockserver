FROM golang:1.15.6-alpine3.12 as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64


WORKDIR /go/src/mockserver

# Copy and download dependencies, these should change
# often so we can cache this layer
COPY go.* ./
RUN go mod download

# Copy source files and build go binary
COPY . .
RUN go build -o output/mockserver cmd/mockserver/main.go

# build final runtime image
FROM scratch
COPY --from=builder /go/src/mockserver/output/mockserver /opt/mockserver

ENTRYPOINT ["opt/mockserver"]