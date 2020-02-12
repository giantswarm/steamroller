FROM golang:1.13-alpine3.10 AS build

WORKDIR /build

# Compile our steamroller executable
ADD main.go .
RUN go build -o steamroller

# ---------

FROM alpine:3.10

WORKDIR /

# Copy the steamroller executable from the build phase
COPY --from=build /build/steamroller /

ENTRYPOINT ["/steamroller"]
