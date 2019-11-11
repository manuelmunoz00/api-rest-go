FROM golang:1.11-alpine AS builder

########
# Prep
########

# add the source
COPY . /go/src/api-rest-go
WORKDIR /go/src/api-rest-go/

########
# Build Go Wrapper
########

# Install go dependencies
RUN apk update && apk add git && go get github.com/gorilla/mux

#build the go app
RUN GOOS=linux GOARCH=amd64 go build -o ./api-rest-go ./main.go

########
# Package into runtime image
########
FROM alpine

RUN apk --no-cache add ca-certificates

# copy the executable from the builder image
COPY --from=builder /go/src/api-rest-go .

ENTRYPOINT ["/api-rest-go"]

EXPOSE 8080