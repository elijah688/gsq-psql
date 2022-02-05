FROM golang:rc-alpine

WORKDIR /gql

RUN apk add jq curl --no-cache

COPY . .

RUN go build gql

ENTRYPOINT [ "./gql" ]
