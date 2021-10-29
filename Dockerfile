FROM golang:1.13-alpine
COPY . .
RUN GOOS=linux go build -o ./app .

FROM alpine
COPY --from=0 ./app /app
ENTRYPOINT /app
