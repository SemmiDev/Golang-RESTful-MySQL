# API SPEC
*Sample Go-Lang RESTful Project*

## GET ALL STUDENTS

Request : 

**GET http://localhost:9090/students
Accept: application/json**

Response :

```json
[
  {
    "id": 5,
    "identifier": "2003113950",
    "name": "sammi3",
    "email": "sammi3@gmail.com",
    "semester": 2,
    "created_at": "2020-03-06T07:04:05Z",
    "updated_at": "2020-03-06T07:04:05Z"
  },
  {
    "id": 4,
    "identifier": "2003113949",
    "name": "sammi2",
    "email": "sammi2@gmail.com",
    "semester": 2,
    "created_at": "2020-03-06T07:04:05Z",
    "updated_at": "2020-03-06T07:04:05Z"
  },
  {
    "id": 3,
    "identifier": "2003113948",
    "name": "sammi",
    "email": "sammi@gmail.com",
    "semester": 2,
    "created_at": "2020-03-06T07:04:05Z",
    "updated_at": "2020-03-06T07:04:05Z"
  }
]
```

## CREATE STUDENT

Request :

**POST http://localhost:9090/students/create
Content-Type: application/json
Accept: application/json**

```json
{
  "identifier": "2023113920",
  "name": "sam3",
  "email": "sam21@gmail.com",
  "semester": 7,
  "created_at": "2020-04-06T07:04:05Z",
  "updated_at": "2020-04-06T07:04:05Z"
}
```
Response :

```json
{
  "status": "CREATED"
}

```

## UPDATE STUDENT

Request :

**PUT http://localhost:9090/students/update
Content-Type: application/json
Accept: application/json**

```json
{
  "id": 10,
  "identifier": "3023113820",
  "name": "sam3up",
  "email": "sam21up@gmail.com",
  "semester": 5,
  "created_at": "2021-04-06T08:01:26Z",
  "updated_at": "2021-04-06T08:01:26Z"
}
```

Response :

```json
{
    "id": 10,
    "identifier": "3023113820",
    "name": "sam3up",
    "email": "sam21up@gmail.com",
    "semester": 5,
    "created_at": "2021-03-06T08:11:58Z",
    "updated_at": "2021-03-06T08:14:12Z"
}
```

## Delete Student
Request :

**DELETE http://localhost:9090/students/delete?id=9
Accept: application/json**

Response :

```json
{
  "status": "DELETED"
}
```

**SAMMIDEV**
**2021**