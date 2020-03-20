FROM golang:1.14-alpine AS build-env

COPY . /cleanarchitecture-src
WORKDIR /cleanarchitecture-src/cmd/cleanarchitecture-sample

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo -o cleanarchitecture-sample \
 && mv ./cleanarchitecture-sample /


FROM scratch

COPY --from=build-env /cleanarchitecture-sample /usr/local/bin/

EXPOSE 8080

CMD ["cleanarchitecture-sample"]
