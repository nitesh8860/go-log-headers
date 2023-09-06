# go-log-headers

**go-log-headers** is a straightforward Golang application designed to log HTTP request headers in JSON format to the standard output (stdout).

## Overview

This application is a useful tool for inspecting incoming HTTP request headers. It can be run both directly on your local machine or within a Docker container.

## Purpose

The primary purpose of this project is to test Ingress controller settings for pushing GeoIP headers into the logs. The following proxy headers are set by the Ingress controller:

```yaml
proxySetHeaders:
  X-geoip-area-code: $geoip_area_code
  X-geoip-city-continent-code: $geoip_city_continent_code
  X-geoip-city-country-code: $geoip_city_country_code
  X-geoip-city-country-code3: $geoip_city_country_code3
  X-geoip-city-country-name: $geoip_city_country_name
  X-geoip-dma-code: $geoip_dma_code
  X-geoip-latitude: $geoip_latitude
  X-geoip-longitude: $geoip_longitude
  X-geoip-region: $geoip_region
  X-geoip-region-name: $geoip_region_name
  X-geoip-city: $geoip_city
  X-geoip-postal-code: $geoip_postal_code
config:
  use-geoip2: true
```

## Local Usage

### Build the Application

To build the application locally, use the following command:

```shell
go build
```

### Run the Application

After building, execute the application as follows:

```shell
./headers
```

### Send a Request

You can send a request to the locally running application using `curl`:

```shell
curl localhost:8082
```

## Docker Usage

### Build a Docker Image

To run the application within a Docker container, first build a Docker image:

```shell
docker build -t go-headers:1.0 .
```

### Run the Docker Container

After building the Docker image, create and run a container:

```shell
docker run --name go-headers go-headers:1.0
```

### Send a Request to the Dockerized Application

You can send a request to the application running inside the Docker container using `curl`:

```shell
curl localhost:8082
```

## Example Output

The application will log the incoming request headers in JSON format to the console. Here's an example of the output:

```json
{
  "Host": "localhost:8082",
  "User-Agent": "curl/7.68.0",
  "Accept": "*/*"
}
```

Feel free to use this tool to inspect HTTP request headers for debugging or analysis purposes.

## License

This project is licensed under the [MIT License](LICENSE).
