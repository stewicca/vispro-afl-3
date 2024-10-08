## AFL - 3

### Table Database
```postgresql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
id         SERIAL PRIMARY KEY,
title      VARCHAR(128) NOT NULL,
artist     VARCHAR(255) NOT NULL,
price      DECIMAL(5,2) NOT NULL
);

INSERT INTO album (title, artist, price)
VALUES
('Blue Train', 'John Coltrane', 56.99),
('Giant Steps', 'John Coltrane', 63.99),
('Jeru', 'Gerry Mulligan', 17.99),
('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

### Base URL

```
http://localhost:8080
```

### Endpoints
### 1. Get All Albums
```
url: /albums
method: GET
Response:
    Status: 200 OK
    Body:
    [
      {
        "id": 1,
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
      },
      {
        "id": 2,
        "title": "Giant Steps",
        "artist": "John Coltrane",
        "price": 63.99
      }
    ]
```
### 2. Get Album By Id
```
url: /albums/:id
method: GET
url params:
    id (required): The ID of the album
response:
    status: 200 OK
    body: 
    {
      "id": 1,
      "title": "Blue Train",
      "artist": "John Coltrane",
      "price": 56.99
    }
    status: 400 Not Found
    body:
    {
      "message": "album not found"
    }
```
### 3. Create a New Album
```
url: /albums
method: POST
request body:
    content-type: application/json
    body:
    {
      "title": "Kind of Blue",
      "artist": "Miles Davis",
      "price": 39.99
    }
response:
    status: 201 Created
    body:
    {
      "id": 5,
      "title": "Kind of Blue",
      "artist": "Miles Davis",
      "price": 39.99
    }
    status: 400 Bad Request
    body:
    {
      "error": "Invalid request body"
    }
```
### 4. Update an Album
```
url: /albums/:id
method: PUT
url params:
    id (required): The ID of the album
request body:
    content-type: application/json
    body:
    {
      "title": "Kind of Blue",
      "artist": "Miles Davis",
      "price": 39.99
    }
response:
    status: 20O OK
    body:
    {
      "message": "album updated successfully"
    }
    status: 404 Not Found
    body:
    {
      "message": "album not found"
    }
```
### 5. Delete an Album
```
url: /albums/:id
method: DELETE
url params:
    id (required): The ID of the album
response:
    status: 20O OK
    body:
    {
      "message": "album deleted successfully"
    }
    status: 404 Not Found
    body:
    {
      "message": "album not found"
    }
```