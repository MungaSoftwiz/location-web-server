# Location Web Server

Location-web-server is a simple web server that provides weather information and personalized greetings based on the client's IP address and location.

## Features

- Fetches client's IP address and geolocation.
- Retrieves current weather data for the client's location.
- Caches geolocation and weather data to reduce API calls.
- Exposes an API endpoint for personalized greetings.

## API Endpoint

### GET `/api/hello`

#### Query Parameters:
- `visitor_name` (optional): The name of the visitor (default: "Guest").

#### Response:
```json
{
  "client_ip": "127.0.0.1",
  "location": "New York",
  "greeting": "Hello, Mark!, the temperature is 11 degrees Celsius in New York"
}

