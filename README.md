# API Shorten URL

### **1. General Structure**
- The `apiproject` package contains the implementation of an HTTP server using the **Chi** framework (`github.com/go-chi/chi/v5`), which is lightweight and well-suited for APIs.
- It includes middleware functionality (like `Logger`, `Recoverer`, `RequestID`), JSON serialization, HTTP redirection, and random code generation.

---

### **2. `sendJson` Function**
This function encapsulates the sending of JSON responses. 

#### Details:
- Sets the `Content-Type` header to `application/json`.
- Converts the response (`Respose`) into JSON using `json.Marshal`.
- Handles errors for serialization and writing to the response.
- Centralizes JSON response handling for cleaner and reusable code.

---

### **3. `NewHandler` Function**
This function sets up the HTTP router (using Chi) and defines the API routes.

#### Configured Routes:
- **`POST /api/shorten`**: Generates a short code for a URL sent in the request body.
- **`GET /{code}`**: Redirects the user to the URL corresponding to the provided code.

#### Middleware Usage:
- **`middleware.Recoverer`**: Recovers from panics in the server to avoid crashes.
- **`middleware.RequestID`**: Generates a unique ID for each request (useful for debugging).
- **`middleware.Logger`**: Logs information for each request handled by the server.

---

### **4. Custom Types**
- **`PostBody`**: Represents the body expected in the `POST /api/shorten` request.
  - Contains the `URL` (string) field sent by the client.
- **`Respose`**: Represents the API responses.
  - Includes `Error` (error message) and `Data` (response data, can be any type).

---

### **5. `genCode` Function**
Generates a random 8-character code used as a short identifier for URLs.

#### Details:
- Uses a set of characters (`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`).
- Seeds the random number generator with the current time (`rand.Seed(time.Now().UnixNano())`) to ensure different codes on each run.

---

### **6. `handlePost` and `handleGet` Functions**
These functions implement the behavior of the API routes.

#### `handlePost(db map[string]string)`
- Decodes the JSON body of the request to get the URL provided by the client.
- Validates the URL using `url.Parse`.
- Generates a short code with `genCode` and stores the URL in the database (`db`).
- Returns the short code as a response.

#### `handleGet(db map[string]string)`
- Retrieves the code from the route parameters.
- Checks if the code exists in the database (`db`).
- If found, redirects the client to the original URL.
- Otherwise, returns a `404 Not Found` error.

### 7. Database (map[string]string)

The db is a map (map[string]string) passed as a parameter to NewHandler and used to store shortened URLs in memory.

Maps short codes (string) to full URLs (string).
Limitation: Being in-memory, the URLs are lost when the server restarts. In production, this would be replaced by a persistent database.