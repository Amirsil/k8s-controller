FROM golang:1.13-alpine
ADD . /
WORKDIR /
RUN GOOS=linux GOFLAGS=-mod=vendor go build -o ./app .

FROM alpine
COPY --from=0 ./app /app
RUN chmod a+x /app
ENTRYPOINT /app
