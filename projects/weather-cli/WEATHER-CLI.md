# Weather CLI

Goal: Create a tool to fetch and display current weather data for a given city.

Features:
Use a public weather API (like OpenWeatherMap).
Display temperature, humidity, wind speed, etc.
Option to get the forecast for multiple days.

Learning: Working with APIs, handling JSON data, Go HTTP package.

Initialize the Project:
```bash
mkdir weather-cli
cd weather-cli
go mod init weather-cli
```

Build the CLI Tool:
```bash
go build -o weather-cli
```

Run the CLI Tool:
```bash
./weather-cli -city "New York"
```