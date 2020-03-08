FROM golang:buster

WORKDIR /app
COPY . /app/
RUN go build -o gql .

ENTRYPOINT ["/app/entrypoint.sh"]
