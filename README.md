# Sales App API 

Given that you already know how to run a buffalo App, I will skip that part.
This README is to let you know what is the request/response format and the endpoints this App is serviing.

## Request example

`GET /clients/`

    http://localhost:3000/clients

## Response example
    {
        "status": 200,
        "data": [
            {
            "id": "b5ca0e83-1aa8-4c98-84ac-20ac52f3a2c1",
            "name": "Avengers",
            "phone_number": "3019995868",
            "rep": "Tony Stark",
            "created_at": "2022-12-02T21:06:03.791659Z",
            "updated_at": "2022-12-02T21:06:03.791659Z"
            }
        ],
        "message": "Clients list"
    }


## Endpoints list

#### Teams
#### `GET /teams`
#### `GET /teams/{id}`
#### `GET /teams/{id}/employees`
#### `POST /teams/create`
#### `DELETE /teams/{id}`

#### Empoyees
#### `GET /employees`
#### `GET /employees/{id}`
#### `POST /employees/create`
#### `DELETE /employees/{id}`

#### Clients
#### `GET /clients`
#### `GET /clients/{id}`
#### `GET /clients/{id}/offers`
#### `POST /clients/create`
#### `DELETE /clients/{id}`

#### Offers
#### `GET /offers`
#### `GET /offers/{id}`
#### `POST /offers/create`
#### `DELETE /offers/{id}`

#### Sales
#### `GET /sales`
#### `GET /sales/{id}`
#### `POST /sales/create`
#### `DELETE /sales/{id}`
