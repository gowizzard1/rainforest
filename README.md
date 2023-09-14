# Rainforest Challenge Resolver

This Rainforest Challenge Resolver is written in Go. The application is designed to make an HTTP GET request to a specified URL, following subsequent URLs provided in the JSON response until it retrieves a `message` property. The application employs goroutines to manage concurrency.

## Prerequisites

1. Go (1.15 or later recommended)
2. Git

## Setup

1. Clone the repository to your local machine.

```bash
git clone https://github.com/gowizzard1/rainforest.git
```

2. Navigate to the project directory.

```bash
cd rainforest
```

## Running the Application

Execute the following command in the project directory:

```bash
go run main.go
```

This will start the application and it will begin making requests to the challenge URL. Once it finds the final message, it will print it to the console.

## Understanding the Code

- `main.go`: Contains the primary logic for the application.
- `ChallengeResponse`: A struct that represents the expected structure of the JSON response.
- `getQueryParam`: A helper function to extract specific query parameters from URLs.
- `makeRequest`: This function makes the request, parses the response, and either follows the next URL or returns the final message. It uses goroutines for concurrency.

## Notes

- Error handling is done through panics for simplicity. In a production environment, errors would be handled more gracefully.
- The application assumes that there will always be a `message` property at the end of the sequence of requests.
