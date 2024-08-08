FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata bash curl nano git cron

COPY --from=builder /app/main /usr/local/bin/main
COPY crontab /etc/crontabs/root

# Set environment variables 
ENV TWILIO_ACCOUNT_SID=""
ENV TWILIO_AUTH_TOKEN=""
ENV TWILIO_IP_LIST_SID=""
ENV TWILIO_ORIGINATION_SID=""
ENV CLOUDFLARE_API_KEY=""
ENV GET_IP_QUERY_URL="https://api.ipify.org?format=text"
ENV DOMAIN_NAME=""

RUN touch /var/log/cron.log

CMD ["crond", "-f", "-l", "2"]
