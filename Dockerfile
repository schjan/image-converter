FROM golang:1.23 AS build-env

# Add namespace here to resolve /vendor dependencies etc.
ENV NAMESPACE github.com/schjan/image-converter
ENV MAINFILE cmd/image-converter/main.go
WORKDIR /go/src/$NAMESPACE

ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags '-w -s'  -a -installsuffix cgo -o /application $MAINFILE

FROM scratch
COPY --from=build-env /application /
ENTRYPOINT [ "/application" ]
