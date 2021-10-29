FROM golang:1.13-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN GOOS=linux GOFLAGS=-mod=vendor go build -o ./app .

FROM alpine
COPY --from=0 ./app /app
ENTRYPOINT /app
