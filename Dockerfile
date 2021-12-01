FROM golang:1.16-alpine

WORKDIR /app

COPY . .

ENTRYPOINT [ "go", "run"]

CMD [ "main.go" ]