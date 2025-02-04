# Receipt Processor

### Running the Server
To run the server, use the following command:

```bash
go run main.go
```

### Running Unit Tests

```bash
go test ./...
```
*Note: This will exclude all tests tagged as integration tests.*

#### Running Integration Tests
To execute integration tests:

```bash
go test -tags=integration ./...
```

*Note: The server must be running before starting integration tests.*

---

### Code Structure
The models for this project have been auto-generated using `oapi-codegen`. It can be regenerated from the spec using `tools.go`
The codebase is divided into separate packages, and unit tests are available for nearly all of them.

---

### Performance Considerations
For improved performance, it might have been better to calculate the points asynchronously when the receipt is initially stored. However, since this was not explicitly required, the calculation currently occurs during read operations.

- **Advantages of current approach**: Calculation rules can change without needing a data modification across all stored entries.
- **Trade-offs**: Storing points during receipt storage may yield better runtime performance but would make rule changes more complicated.

---

### Error Handling
Validation errors are surfaced to clients with a **generic error response**.  
This approach is generally considered a best practice as it avoids exposing internal implementation details.

If the API was to grow and more endpoints added, I would consider introducing a validation library (for example ozzo)

---

### Additional Notes
Note: This solution does not include considerations for API security.      

Due to new pricing requirements for Postman, no Postman collection is included at this time.

For production readiness, a logging library should be added, instead of using `fmt.Printf`
