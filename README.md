# Advanced Programming 1 Assignment  |  Erkinkyzy Bakyt
## Making http server with post request

### Checking POST method
> Method checks if the request method is POST
```go
  if r.Method != http.MethodPost {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }
```

### Success request
> Request body (input):
```json
{
  "username": "erk1nqz",
  "email": "someEmail@emailService.com",
  "password": "somePassword"
}
```

> Response body (output/result):
```json
{
	"status": "success",
	"message": "User erk1nqz successfully registered with password - somePassword"
}
```

### Unsuccessful request 1
> Request body (input):
```json
{
  "username": "erk1nqz",
  "password": "somePassword"
}
```

> Response body (output/result):
```json
Invalid JSON message
```

### Unsuccessful request 2
> Request body (input):
```json
{
  username: "erk1nqz"
}
```

> Response body (output/result):
```json
Invalid JSON format
```
