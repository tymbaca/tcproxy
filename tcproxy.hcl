group {
	port = 8080
    protocol = "tcp"
    strategy = "round_robin"

    // middlewares = [
    //     "logger"
    // ]

    // target { addr = "localhost:8090" }
    target { addr = "localhost:8091" }
    target { addr = "localhost:8092" }
}
