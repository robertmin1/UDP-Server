# UDP-Server

The project is a simple UDP server where data is echoed to the server using a bash script.

# Building

#### Prerequisites:

1. Ensure you have the Go tools installed.

#### Option A : Using Go build commands

1. Clone `UDP-Server` via Git

2. Run `go run UDP-server.go` to start up the server

4. Open a separate terminal and run `./send-to-server.bash` to echo data to the server.

5. The echoed data `hello world` will be displayed in the server terminal

## TODO.
    Fix Cirrus Test
        - The test works for fedora but doesn't work for debian. Error: "./UDP-server.go:34:12: undefined: os.WriteFile"
        - Implement Go Lint Tests.
        - Implement better tests (Effiecent tests)

## Information About The Program

UDP Server

