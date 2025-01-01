group {
	port = 8080
    strategy = "random"

    target { addr = "localhost:8090" }
    target { addr = "localhost:8091" }
    target { addr = "localhost:8092" }
}
