# UDP-Sever

The project is a simple UDP server where data is echoed to the server using a bash script.

# Building

#### Prerequisites:

1. Ensure you have the Go tools installed.

#### Option A : Using Go build commands

1. Clone `UDP-Sever` via Git

2. Run `go run UDP-sever.go` to start up the sever

4. Open a separate terminal and run `./send-to-sever.bash` to echo data to the sever.

5. The echoed data `hello world` will be displayed in the sever terminal

## TODO.
    Fix Cirrus Test
        - The test works for fedora but doesn't work for debian. Error: "./UDP-sever.go:34:12: undefined: os.WriteFile"
        - Implement Go Lint Tests.
        - Implement better tests (Effiecent tests)

## Information About The Program

UDP Sever

