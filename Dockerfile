FROM golang:buster

WORKDIR /app
COPY . /app/

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.8
RUN go build -o gql .

ENTRYPOINT ["/app/entrypoint.sh"]
