# API Project Movies

This project is a simple API in Go that allows you to search for movie information using the OMDB (Open Movie Database) API. The goal of the API is to perform movie searches based on parameters provided by the user, such as the movie title (`s` parameter) and the API key (`apikey`).

## Features

- **Movie Search**: The API allows you to search for movies by providing the movie title as the `s` parameter and an OMDB API key as `apikey`.
- **JSON Responses**: The response will be returned in JSON format, containing the movie data or an error message in case of an issue.

## Endpoints

### `GET /`

This endpoint allows you to search for a movie. The required query parameters are:

- `apikey`: Your OMDB API key.
- `s`: The movie title to search for.

#### Parameters

- **apikey** (required): The OMDB API key.
- **s** (required): The title of the movie to search for.

#### Responses

- **200 OK**: If the request is successful, the response will contain the movie data in JSON format.
  
- **400 Bad Request**: If the `s` (movie title) parameter is not provided.
  
- **401 Unauthorized**: If the `apikey` parameter is missing or invalid.
  
- **502 Bad Gateway**: If there is an error fetching data from the OMDB API.

### Example Request

```bash
curl "http://localhost:8085/?apikey=<API-KEY-HERE>&s=Blade" 
```

### Example Response

#### Success (200 OK):

```json
{
  "Data": {
    "Title": "Blade Runner",
    "Year": "1982",
    "Genre": "Sci-Fi, Thriller",
    "Director": "Ridley Scott",
    "Actors": "Harrison Ford, Rutger Hauer, Sean Young",
    "Plot": "A blade runner must pursue and try to terminate four replicants who stole a ship in space, and have returned to Earth to find their creator."
  }
}
```

#### Missing Parameter (400 Bad Request):

```json
{
  "Error": "search parameter 's' is required"
}
```

#### Missing API Key (401 Unauthorized):

```json
{
  "Error": "apikey parameter is required"
}
```

#### Error Fetching from OMDB (502 Bad Gateway):

```json
{
  "Error": "something wrong with omdb"
}
```

## How to Run the Project

### Prerequisites

- A valid OMDB API key. You can obtain your key from [OMDB API](https://www.omdbapi.com/apikey.aspx).

### Steps to Run Locally

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/frankdias92/go-playground.git
   git switch api-project-movies
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Start the server:

   ```bash
   go run main.go
   ```

4. The API will be available at `http://localhost:8085`. You can test it with `curl`:

   ```bash
   curl "http://localhost:8085/?apikey=<YOU-API-KEY>&s=Blade"
   ```

## Technical Details

### Architecture

- **Chi**: The `chi` package is used for HTTP routing. It provides a simple yet powerful way to build RESTful APIs.
  
- **OMDB API**: The OMDB API is queried to search for movies using the `apikey` and `s` (movie title) parameters.

- **Middleware**:
  - `Recoverer`: Recovers the server from panics and returns a 500 HTTP response in case of an error.
  - `RequestID`: Generates a unique ID for each request.
  - `Logger`: Logs request information.

- **Error Handling and Responses**: The API handles errors in a structured way, returning error messages in JSON format with the appropriate status code.

