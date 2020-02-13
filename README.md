# films-server
Basic go project

## Running locally
```zsh
# Build a local image
$ go build
# Run local image
$ ./go-docker
```

## Building docker image
```zsh
# Build a docker image
$ docker build -t go-docker .
```

## Starting
```zsh
# build image
$ ./go-docker
```

## Stopping
```zsh
# list out docker containers
$ docker container ls
# Find image and remove
$ docker stop <id>
```

#### Resources
- https://www.callicoder.com/docker-golang-image-container-example/