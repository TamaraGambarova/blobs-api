FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/<%= packageName %>
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/<%= serviceName %> /go/src/<%= packageName %>


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/<%= serviceName %> /usr/local/bin/<%= serviceName %>
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["<%= serviceName %>"]
