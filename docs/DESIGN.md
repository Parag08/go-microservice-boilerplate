# Design

As part of the application we will focus on CRUD operation for user. So the core idea would be to develop

## User Structure

```go
type User struct {
    Name string `json:"name" bson:"name" validate:"required"`
    Email string `json:"email" bson:"email" validate:"required,email"`
    Password string `json:"password" bson:"password" validate:"required,passwd"`
}
```

### Create User

### Get User

### Modify User

### Delete User
