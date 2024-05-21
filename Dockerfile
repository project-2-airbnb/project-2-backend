FROM golang:1.21.5-alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o beapi

CMD [ "./beapi" ]
