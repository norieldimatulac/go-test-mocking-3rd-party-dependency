# go-test-mocking-3rd-party-dependency

# main:

Basic example of golang Employee object with unit test.

    To run the code:
        1. docker run --rm --privileged -ti --name go-test-mocking-3rd-party-dependency -v $PWD:/go/src/go-test-mocking-3rd-party-dependency -w /go/src/go-test-mocking-3rd-party-dependency norieldimatulac/alpine-golang-dep:1.0.0 bash
        2. dep ensure
        3. cd main
        4. go run main.go
        5. go test -cover -v

# main2:

Example golang code that is using 3rd party dependency (REDIS). The REDIS connection and client.HGetAll inside GetFullName() function is now problematic in writing unit test. This will require us to have an actual running REDIS database when executing "go test" to run the unit tests.

    To run the code:
        1. Run REDIS docker container
            $ docker run --name some-redis -p 6379:6379 -d redis
        2. Load test data in to redis database
            $ docker exec -ti some-redis bash
            $ redis-cli
            $ hmset emp:001 FirstName "Bronzon" LastName "Yatayan"
            $ hmset emp:002 FirstName "Jun" LastName "Bols"
            $ hmset emp:003 FirstName "Juan" LastName "Dela Cruz"
        3. docker run --rm --privileged -ti --name go-test-mocking-3rd-party-dependency -v $PWD:/go/src/go-test-mocking-3rd-party-dependency -w /go/src/go-test-mocking-3rd-party-dependency norieldimatulac/alpine-golang-dep:1.0.0 bash
        4. dep ensure
        5. cd main2
        6. go run main2.go
        
# main3:

Using dependency injection and mockery (https://github.com/vektra/mockery) to mock or fake the redis object in unit test.

    To run the code:
        1. Run REDIS docker container
            $ docker run --name some-redis -p 6379:6379 -d redis
        2. Load test data in to redis database
            $ docker exec -ti some-redis bash
            $ redis-cli
            $ hmset emp:001 FirstName "Bronzon" LastName "Yatayan"
            $ hmset emp:002 FirstName "Jun" LastName "Bols"
            $ hmset emp:003 FirstName "Juan" LastName "Dela Cruz"
        3. docker run --rm --privileged -ti --name go-test-mocking-3rd-party-dependency -v $PWD:/go/src/go-test-mocking-3rd-party-dependency -w /go/src/go-test-mocking-3rd-party-dependency norieldimatulac/alpine-golang-dep:1.0.0 bash
        4. dep ensure
        5. Installing mockery
            $ go get github.com/vektra/mockery/.../
            $ apk add --update gcc
            $ apk add --update gcc musl-dev
        6. cd main3
        7. $GOPATH/bin/mockery -name Database
        8. go run main3.go database.go redis.go
        9. go test -cover -v
        
To change the database from redis to mongo:

    1. Run a mongo docker container
        $ docker run --name some-mongo -p 27017:27017 -d mongo
    2. Load test data in to mongo database
        $ docker exec -ti some-mongo bash
        $ mongo
        $ use test
        $ db.employee.insert({EmpID:"emp:001", FirstName: "Bronzon", LastName: "Yatayan"})
        $ db.employee.insert({EmpID:"emp:002", FirstName: "Jun", LastName: "Bols"})
        $ db.employee.insert({EmpID:"emp:003", FirstName: "Juan", LastName: "Dela Cruz"})
    3. cd main3
    4. vim database.go
    5. In NewDatabase() function, change redis to mongo r := &Mongo{}
    6. :wq
    7. go run main3.go database.go mongo.go
    8. go test -cover -v
