#TODO

    Static linking:

    Go 1.5:

    go build -ldflags "-extldflags -static"

    With Go 1.6 I had to use:

    go build -ldflags "-linkmode external -extldflags -static"


