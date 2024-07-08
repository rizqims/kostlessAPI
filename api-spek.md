template response:
```go
{
    "code": int,
    "message": "string",
    "data": any
}
```


### Owner

#### POST RegisterOwner

request:
```go
{
    "fullName": "string",
    "username": "string",
    "password": "string",
    "email": "string",
    "phoneNumber": "string"
}
```

response:
```go
{
    "code": int,
    "message": "string",
    "data": {
        "id": "string",
        "fullName": "string",
        "username": "string",
        "password": "string",
        "email": "string",
        "phoneNumber": "string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
}
```

#### POST LoginOwner

request:
```go
{
    "username": "string",
    "password": "string"
}
```

response:
```go
{
    "code": int,
    "message": "string",
    "data": {
        "token":"string"
    }
}
```

### Seeker

#### POST RegisterSeeker

request:
```go
{
    "fullName": "string",
    "username": "string",
    "password": "string",
    "email": "string",
    "phoneNumber": "string",
    "status":"string",
}
```

response:
```go
{
    "code": int,
    "message": "string",
    "data": {
        "id": "string",
        "fullName": "string",
        "username": "string",
        "password": "string",
        "email": "string",
        "phoneNumber": "string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
}
```

### POST LoginSeeker

request:
```go
{
    "username": "string",
    "password": "string"
}
```

response:
```go
{
    "code": int,
    "message": "string",
    "data": {
        "token":"string"
    }
}
```

### Kost

#### POST CreateKost

request:
```go
{
    "name": "string",
    "address": "string",
    "roomCount": int,
    "coordinate": "string",
    "desc": "string",
    "rules": "string"
}
```

response:
```go
    "code": int,
    "message": "string",
    "data": {
        "id": "string",
        "name": "string",
        "address": "string",
        "roomCount": int,
        "coordinate": "string",
        "desc": "string",
        "rules": "string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
```

### Room

#### POST CreateRoom

request:
```go
{
    "name": "string",
    "type": "string",
    "desc": "string",
    "avail": "string",
    "price": INT,
}
```

response:
```go
    "code": int,
    "message": "string",
    "data": {
        "id":"string",
        "name": "string",
        "type": "string",
        "desc": "string",
        "avail": "string",
        "price": INT,
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
```