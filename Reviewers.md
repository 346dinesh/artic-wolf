# For Reviewers

- This implementation uses the  Go web framework "Gin", "context" router for convenient path variables.
- Data is stored in a `map[string]Risk` named `MemoryOfRisks`.
- We generate a UUID upon risk creation using the `github.com/google/uuid` package.
- We enforce valid states by using a simple `map[string]bool`.
- Errors are handled with relevant HTTP status codes:
  - `400 Bad Request` when JSON or state is invalid.
  - `404 Not Found` when no risk is found for a particular ID.
  - `500 Internal Server Error` if JSON encoding fails unexpectedly.
- The service does not use HTTPS or any authentication for this interview assignment. In a production environment, both would be highly recommended.
