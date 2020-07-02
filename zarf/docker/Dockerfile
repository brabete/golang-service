# Build the Go Binary.
FROM golang:1.14.4 as build_sales-api
ENV CGO_ENABLED 0

# Create a location in the container for the source code. Using the
# default GOPATH location.
RUN mkdir -p /service

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
# COPY go.* /service/
# WORKDIR /service
# RUN go mod download

# Copy the source code into the container.
WORKDIR /service
COPY . .

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /service/app/sales-api
RUN go build -ldflags "-X main.build=prod"

# Run the Go Binary in Alpine.
FROM alpine:3.7

COPY --from=build_sales-api /service/private.pem /app/private.pem
COPY --from=build_sales-api /service/app/sales-api/sales-api /app/main
WORKDIR /app
CMD ["/app/main"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="${PACKAGE_NAME}" \
      org.opencontainers.image.authors="William Kennedy <bill@ardanlabs.com>" \
      org.opencontainers.image.source="https://github.com/ardanlabs/service/cmd/${PACKAGE_PREFIX}${PACKAGE_NAME}" \
      org.opencontainers.image.revision="${VCS_REF}" \
      org.opencontainers.image.vendor="Ardan Labs"
