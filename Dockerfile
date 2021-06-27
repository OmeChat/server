FROM golang:1.16 AS builder

WORKDIR /build

COPY . .
RUN go mod tidy

ENV CGO_ENABLED=0

RUN go build ./cmd/main.go

FROM alpine:3 AS final

COPY --from=builder /build/main .

RUN mkdir data
RUN touch ./data/user.json
RUN echo {} > ./data/user.json
RUN touch ./data/client.json
RUN echo {} > ./data/client.json

RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]