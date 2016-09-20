#TODO

    Static linking:

    Go 1.5:

    go build -ldflags "-extldflags -static"

    With Go 1.6 I had to use:

    go build -ldflags "-linkmode external -extldflags -static"


### Reduce file size

    make clean && make docker-scratch

    $ docker login
    Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
    Username (bayugyug): bayugyug
    Password:
    Login Succeeded

    
    $ docker push bayugyug/freeformatter:scratch
    The push refers to a repository [docker.io/bayugyug/freeformatter]
    3485c9ef6f83: Pushed
    scratch: digest: sha256:f3844a86e39d47309e761945958056282865a2abcad9d60585ce2c191db198df size: 528

