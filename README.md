# runners-api :runner:

Developing the API for the Runners Social Network (work in progress). This project aims to enhance my expertise in Golang. Stay tuned for updates! :smile:

## How to run locally

After cloning this repository. You need to have a database configured, I used MySQL and you can run the following SQL commands to create your database.

    CREATE DATABASE IF NOT EXISTS runners;
    USE runners;

    DROP TABLE IF EXISTS followers;
    DROP TABLE IF EXISTS users;

    CREATE TABLE users(
        id int auto_increment primary key,
        name varchar(50) not null,
        nick varchar(50) not null unique,
        email varchar(50) not null unique,
        password varchar(100) not null,
        createOn timestamp default current_timestamp()
    ) ENGINE=INNODB;

    CREATE TABLE followers(
        user_id int not null,
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

        follower_id int not null,
        FOREIGN KEY (follower_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

        primary key (user_id, follower_id)
    ) ENGINE=INNODB;

Then you need to create a `.env` file with the below content:

    DB_USER=yourDatabaseUser
    DB_PASSWORD=yourDatabasePassword
    DB_NAME=yourDatabaseName

    SECRET_KEY=yourSecretKey

Finally you can run the commands below and play :smile:

Download the dependencies from `go.mod`

    go mod download

Running the API:

    go run main.go

## API

### Login's endpoint

### Login Runner User

`POST /login`

    curl -i -X POST -H "Content-Type: application/json" -d '{"email":"email@gmail.com","password":"123"}' http://localhost:5000/login

#### Response

    HTTP/1.1 200 OK
    Date: Wed, 03 Apr 2024 20:50:41 GMT
    Content-Length: 144
    Content-Type: text/plain; charset=utf-8

    {"id":5, "token": "token"}

### User's endpoint

Pay attention to the user token, you need it to make some requests.

### Create an Runner User

`POST /users`

    curl -i -X POST -H "Content-Type: application/json" -d '{"name":"name","nick":"nick","email":"email@gmail.com","password":"123"}' http://localhost:5000/users

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 03 Apr 2024 21:08:10 GMT
    Content-Length: 171

    {"id":5,"name":"name","nick":"nick","email":"email@gmail.com","password":"$2a$10$.k1Oad30NYh3joIR/ShWkOOekfOkiLryl15O4T4XKKQy6Zkh56kBm","createOn":"0001-01-01T00:00:00Z"}

### Search Runner User(s) by name or nick

Here you use query parameters.

`GET /users?user={nameOrNick}`

    curl -i -X GET -H "Authorization: Bearer tokenHere" "http://localhost:5000/users?user=name"

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 03 Apr 2024 21:17:18 GMT
    Content-Length: 104

    [{"id":5,"name":"name","nick":"nick","email":"email@gmail.com","createOn":"2024-04-03T18:08:09-03:00"}]

### Search Runner User by ID

`GET /users/{userID}`

    curl -i -X GET -H "Authorization: Bearer tokenHere" "http://localhost:5000/users/5"

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 03 Apr 2024 21:31:49 GMT
    Content-Length: 102

    {"id":5,"name":"name","nick":"nick","email":"email@gmail.com","createOn":"2024-04-03T18:08:09-03:00"}

### Update Runner User

Here you can update name, nick and email.

`PUT /users/{userID}`

    curl -i -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer tokenHere" -d '{"name":"email2","nick":"email2","email":"email2@gmail.com"}' http://localhost:5000/users/5

#### Response

    HTTP/1.1 204 No Content
    Content-Type: application/json
    Date: Wed, 03 Apr 2024 22:02:01 GMT


### Delete Runner User

`DELETE /users/{userID}`

    curl -i -X DELETE -H "Authorization: Bearer tokenHere" http://localhost:5000/users/5

#### Response

    HTTP/1.1 204 No Content
    Content-Type: application/json
    Date: Wed, 03 Apr 2024 22:08:32 GMT


### Follow Runner User

The `userID` is the user to be followed.

`POST /users/{userID}/follow`

    curl -i -X POST -H "Authorization: Bearer tokenHere" http://localhost:5000/users/2/follow

#### Response

    HTTP/1.1 204 No Content
    Content-Type: application/json
    Date: Thu, 04 Apr 2024 20:44:14 GMT

### Unfollow Runner User

The `userID` is the user to be unfollowed.

`POST /users/{userID}/unfollow`

    curl -i -X POST -H "Authorization: Bearer tokenHere" http://localhost:5000/users/2/unfollow

#### Response

    HTTP/1.1 204 No Content
    Content-Type: application/json
    Date: Fri, 05 Apr 2024 18:27:22 GMT

### Search for Followers

The `userID` is the user who wants to see his followers.

`GET /users/{userID}/followers`

    curl -i -X GET -H "Authorization: Bearer tokenHere" http://localhost:5000/users/1/followers

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Thu, 11 Apr 2024 18:34:49 GMT
    Content-Length: 212

    [{"id":2,"name":"User 2","nick":"user2","email":"user2@gmail.com","createOn":"2024-04-04T15:35:00-03:00"},
    {"id":3,"name":"User 3","nick":"user3","email":"user3@gmail.com","createOn":"2024-04-04T15:35:00-03:00"}]

### Search for Following

The `userID` is the user who wants to see who he is following.

`GET /users/{userID}/following`

    curl -i -X GET -H "Authorization: Bearer tokenHere" http://localhost:5000/users/2/following

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Thu, 11 Apr 2024 19:21:55 GMT
    Content-Length: 107

    [{"id":1,"name":"User 1","nick":"user1","email":"user1@gmail.com","createOn":"2024-04-04T15:35:00-03:00"}]
    