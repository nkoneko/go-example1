# syntax=docker/dockerfile:1
FROM arm64v8/golang:1.17.3 AS build
WORKDIR /example
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGI_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags netgo -ldflags '-w' -a -installsuffix cgo -o app

FROM gcr.io/distroless/static:nonroot-arm64
COPY --from=build /example/app /app
CMD ["/app"]
