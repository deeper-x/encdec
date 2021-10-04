```securefile``` is a simple golang library for encrypting and decrypting files.

## Installation

To install the package, run:
```bash
go get -v github.com/deeper-x/securefile
```

## Usage

To encrypt data, run:
```go
import "github.com/deeper-x/securefile"

key := "<YOUR PASSWD>"
_, err := securefile.Encrypt("/path/to/resource", []byte(key))
if err != nil {
	panic(err)
}
```

To decrypt data, run:
```go
import "github.com/deeper-x/securefile"
            
key := "<YOUR PASSWD>"
_, err := securefile.Decrypt("/path/to/resource", []byte(key))
if err != nil {
	panic(err)
}
```

## Test
```go
go test -v ./...
```

## License
[MIT License](https://opensource.org/licenses/MIT)
