FROM golang:1.15.6 as builder

# Add the keys
ARG GIT_DOMAIN
ENV GIT_DOMAIN=$GIT_DOMAIN

ARG GIT_USERNAME
ENV GIT_USERNAME=$GIT_USERNAME

ARG GIT_PASSWORD
ENV GIT_PASSWORD=$GIT_PASSWORD

RUN git config --global url."https://${GIT_USERNAME}:${GIT_PASSWORD}@${GIT_DOMAIN}/".insteadOf "https://${GIT_DOMAIN}/"

WORKDIR /app
COPY . /app
RUN make build


# APP
FROM alpine:latest
# mailcap adds mime detection and ca-certificates help with TLS (basic stuff)
RUN apk --no-cache add tzdata ca-certificates mailcap && addgroup -S app && adduser -S app -G app
RUN echo "Asia/Dhaka" > /etc/timezone
RUN cp /usr/share/zoneinfo/Asia/Dhaka /etc/localtime

USER app
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/bin/intkey /app/intkey
CMD ["intkey"]
