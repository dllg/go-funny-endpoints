![test results](https://github.com/dllg/go-funny-endpoints/workflows/Test/badge.svg)

# Funny endpoints

Simple go-application that interacts with different open api's to create text that you can receive on GET endpoints. See [swagger.yml](swagger.yml).

It uses [gin web framework](https://github.com/gin-gonic/gin) to handle http requests. For logging it uses [logrus](https://github.com/sirupsen/logrus).

## Testing
Testing of [funny.go](./funny/funny.go) is done using dependency injection. Mocking of [httpclient](./httpclient/httpclient.go) is done using [mockgen](https://github.com/golang/mock).

Testing of [router.go](./router/router.go) is done using package httptest.

## Building

### For local computer
Use make to build the application:
```bash
make build
```

### For docker
Use make to build the application in docker:
```bash
make docker-build
```

## Running
### Locally
The application will be put in the build folder. To run it invoke:
```bash
export PORT="18080"
./build/go-funny-endpoints
```

### In docker
Use docker-compose to start the application:
```bash
docker-compose up -d
```

## Examples of using it
### Get an advice
```bash
$ curl http://localhost:18080/v1/advice
{"message":"If you don't want something to be public, don't post it on the Internet."}
```

### Get a Chuck Norris Joke
```bash
$ curl http://localhost:18080/v1/chucknorris
{"message":"there is no use crying over spilled milk, unless its Chuck Norris' milk because then your gonna die"}
```

### Get a dad joke
```bash
$ curl http://localhost:18080/v1/dadjoke
{"message":"What do you call an eagle who can play the piano? Talonted!"}
```

### Get a random message
```bash
$ curl http://localhost:18080/v1/random
{"message":"Advice: When painting a room, preparation is key. The actual painting should account for about 40% of the work."}
```
