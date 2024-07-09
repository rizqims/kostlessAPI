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

#### PUT UpdateProfile

path variable: `:id`

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

#### GET GetOwner

path variable: `:id`

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

#### PUT UpdateProfile

path variable: `:id`

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
        "status":"string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
}
```

#### DELETE DeleteSeeker

path variable: `:id`

response:
```go
{
    "code": int,
    "message": "string",
    "data": 0
}
```

#### GET GetSeekerByID

path variable: `:id`

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
        "status":"string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
}
```

#### GET GetAllSeeker

response:
```go
{
    "code": int,
    "message": "string",
    "data": [
        {
        "id": "string",
        "fullName": "string",
        "username": "string",
        "password": "string",
        "email": "string",
        "phoneNumber": "string",
        "status":"string",
        "attitudePoints": int,
        "roomID":"string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    },
    {
        "id": "string",
        "fullName": "string",
        "username": "string",
        "password": "string",
        "email": "string",
        "phoneNumber": "string",
        "status":"string",
        "attitudePoints": int,
        "roomID":"string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
    ]
}
```

#### POST UpdateAp

request:
```go
{
    "id":"string",
    "attitudePoints": int,
}
```

if ap is lower than 5, send an email or something to notify seekers. if it has more than 10 points, notify the owners instead to add vouchers to seekers

response:
```go
{
    "code": int,
    "message": "string",
    "data": 0
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
{
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
}
```

#### PUT UpdateKost

path variable: `:id`

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
{
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
}
```

#### DELETE DeleteKostByID

path variable: `:id`

response:
```go
{
    "code": int,
    "message": "string",
    "data": 0
}
```

#### GET GetDetailedKost

response:
```go
{
    "code": int,
    "message": "string",
    "data": {
        "id": "string",
        "name": "string",
        "address": "string",
        "roomCount": int,
        "rooms":[
            {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
            },
            {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
            }
        ]
        "coordinate": "string",
        "desc": "string",
        "rules": "string",
        "createdAt": time.Time,
        "updatedAt": time.Time
    }
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
{
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
}
```

#### GET GetAll

response:
```go
{
    "code": int,
    "message": "string",
    "data": [
        {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        },
        {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        }
    ]
}
```

#### GET GetRoomByID

path variable: `:id`

response:
```go
{
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
}
```

#### GET GetByAvail

path variable: `:avail`

response:
```go
{
    "code": int,
    "message": "string",
    "data": [
        {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        },
        {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        }
    ]
}
```

#### GET GetRoomByBudget

query param: `?budget=500000`

desc: example, if seeker input 500000, display room that has the price of idr 50000 or lower

response:
```go
{
    "code": int,
    "message": "string",
    "data": [
        {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        },
        {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        }
    ]
}
```

#### PUT UpdateRoom

path variable: `:id`

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
{
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
}
```

#### DELETE DeleteRoom

path variable: `:id`

response:
```go
{
    "code": int,
    "message": "string",
    "data": 0
}
```

### Trans

#### POST CreateTrans

request:
```go
{
    "roomID":"string",
    "seekerID":"string",
    "startDate": time.Time,
    "endDate": time.Time,
    "total": int,
    "payLater": bool,
    "dueDate": time.Time
}
```

response:
```go
{
    "code": int,
    "message": "string",
    "data":{
        "id":"string",
        "roomID": "string",
        "seekerID": "string",
        "startDate": time.Time,
        "endDate": time.Time,
        "total": int,
        "payLater": bool,
        "dueDate": time.Time
    }
}
```

#### GET GetTransByID

path variable: `:id`

response:
```go
{
    "code": int,
    "message": "string",
    "data":{
        "id":"string",
        "room": {
            "id":"string",
            "name": "string",
            "type": "string",
            "desc": "string",
            "avail": "string",
            "price": INT,
            "createdAt": time.Time,
            "updatedAt": time.Time
        },
        "seeker": {
            "id": "string",
            "fullName": "string",
            "username": "string",
            "password": "string",
            "email": "string",
            "phoneNumber": "string",
            "createdAt": time.Time,
            "updatedAt": time.Time
        },
        "startDate": time.Time,
        "endDate": time.Time,
        "total": int,
        "payLater": bool,
        "dueDate": time.Time
    }
}
```

#### GET GetPaylaterList


response:
```go
{
    "code": int,
    "message": "string",
    "data":[
            {
            "id":"string",
            "room": {
                "id":"string",
                "name": "string",
                "type": "string",
                "desc": "string",
                "avail": "string",
                "price": INT,
                "createdAt": time.Time,
                "updatedAt": time.Time
            },
            "seeker": {
                "id": "string",
                "fullName": "string",
                "username": "string",
                "password": "string",
                "email": "string",
                "phoneNumber": "string",
                "createdAt": time.Time,
                "updatedAt": time.Time
            },
            "startDate": time.Time,
            "endDate": time.Time,
            "total": int,
            "payLater": bool,
            "dueDate": time.Time
        }
    ]
}
```