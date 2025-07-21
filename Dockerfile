FROM golang:1.24.5-alpine AS builder

WORKDIR /dist

COPY . .

RUN go build -o processor ./...

FROM ghcr.io/mia-platform/integration-connector-agent

LABEL maintainer="%CUSTOM_PLUGIN_CREATOR_USERNAME%" \
  name="mia_template_service_name_placeholder" \
  description="%CUSTOM_PLUGIN_SERVICE_DESCRIPTION%" \
  eu.mia-platform.url="https://www.mia-platform.eu"

ENV SERVICE_VERSION="0.0.0"

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
# Import the certs from the builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER 1000:10000

COPY --from=builder /dist/processor /var/run/processor
