FROM golang:alpine as base

WORKDIR /go/src/github.com/jamesRaybould/concourseVisualiser

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM base as build

RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

FROM scratch

COPY --from=build /go/bin/concourseVisualiser /concourseVisualiser
