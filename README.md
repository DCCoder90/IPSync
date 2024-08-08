# IPSync

This project is designed to periodically check the current public IP address and update DNS records in Cloudflare and IP access lists in Twilio. It uses a Go application to fetch the IP address, and a Docker container to run the application with cron.  This is primarily due to my ISP not offering a static IP Address and me needing some way to keep my self-hosted services updated along with my IP Phone.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Docker Setup](#docker-setup)
- [Contributing](#contributing)
- [License](#license)

## Installation

To build and run this project locally, ensure you have Go and Docker installed on your machine.

### Clone the repository

```sh
git clone https://github.com/yourusername/ip-updater.git
cd IPSync


### Build the Go application

```sh
go build -o main .
```

### Run the application

```sh
./main
```

## Usage

The application checks the public IP address and updates it in Cloudflare and Twilio services. It is designed to be run periodically using cron.

## Environment Variables

Ensure the following environment variables are set:

- `TWILIO_ACCOUNT_SID`: Your Twilio account SID.
- `TWILIO_AUTH_TOKEN`: Your Twilio auth token.
- `TWILIO_IP_LIST_SID`: Twilio IP access list SID.
- `TWILIO_ORIGINATION_SID`: Twilio origination SID.
- `CLOUDFLARE_API_KEY`: Your Cloudflare API key.
- `GET_IP_QUERY_URL`: URL to query the current public IP address (default: `https://api.ipify.org?format=text`).

## Docker Setup

### Dockerfile

The Dockerfile is configured to build the Go application and set up a cron job to run it periodically.

### Build the Docker image

```sh
docker build -t IPSync .
```

### Run the Docker container

```sh
docker run -d --name IPSync -e TWILIO_ACCOUNT_SID=your_sid -e TWILIO_AUTH_TOKEN=your_auth_token -e TWILIO_IP_LIST_SID=your_ip_list_sid -e TWILIO_ORIGINATION_SID=your_origination_sid -e CLOUDFLARE_API_KEY=your_cloudflare_api_key -e GET_IP_QUERY_URL=https://api.ipify.org?format=text IPSync
```

### Docker Compose

You can also use Docker Compose to set up and run the container. Ensure you have a `docker-compose.yml` file with the following content:

```yaml
version: '3.8'

services:
  ip-updater:
    build: .
    environment:
      - TWILIO_ACCOUNT_SID=your_sid
      - TWILIO_AUTH_TOKEN=your_auth_token
      - TWILIO_IP_LIST_SID=your_ip_list_sid
      - TWILIO_ORIGINATION_SID=your_origination_sid
      - CLOUDFLARE_API_KEY=your_cloudflare_api_key
      - GET_IP_QUERY_URL=https://api.ipify.org?format=text
```

Run the Docker Compose setup:

```sh
docker-compose up -d
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.