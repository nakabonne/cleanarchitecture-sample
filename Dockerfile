FROM golang:1.11-alpine AS build-env

COPY . ${GOPATH}/src/github.com/nakabonne/cleanarchitecture-sample
WORKDIR ${GOPATH}/src/github.com/nakabonne/cleanarchitecture-sample

RUN cd src/cleanarchitecture-sample; \
    GOPATH=${GOPATH}/src/github.com/nakabonne/cleanarchitecture-sample:$GOPATH \
    GOPATH=${GOPATH}/src/github.com/nakabonne/cleanarchitecture-sample/vendor:$GOPATH \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo -o cleanarchitecture-sample \
 && mv ./cleanarchitecture-sample /


FROM scratch

COPY --from=build-env /cleanarchitecture-sample /usr/local/bin/

EXPOSE 8080

CMD ["cleanarchitecture-sample"]
