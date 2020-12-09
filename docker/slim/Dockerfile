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


FROM alpine:latest as certs
RUN apk --update add ca-certificates tzdata
RUN echo "Asia/Dhaka" > /etc/timezone
RUN cp /usr/share/zoneinfo/Asia/Dhaka /etc/localtime
RUN addgroup -S app && adduser -S app -G app


# APP
FROM scratch as app
USER app
WORKDIR /app
EXPOSE 8080
COPY --from=certs /etc/timezone /etc/timezone
COPY --from=certs /etc/localtime /etc/localtime
COPY --from=certs /etc/passwd /etc/passwd
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/bin/intkey /usr/local/bin/intkey
CMD ["intkey"]
