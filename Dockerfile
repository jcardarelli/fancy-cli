FROM arm64v8/golang:1.21 AS build
WORKDIR /app
RUN mkdir cmd database types vendor
COPY cmd/ ./cmd/
COPY database/ ./database/
COPY types/ ./types/
COPY vendor/ ./vendor/
COPY go.mod go.sum main.go ./
RUN CGO_ENABLED=1 go build -mod vendor -installsuffix cgo -o ./fancy

FROM debian:12-slim
WORKDIR /go
COPY --from=build /app/fancy .
ENTRYPOINT ["/go/fancy"]
CMD ["restaurant"]
