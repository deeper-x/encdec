```encdec```` is a simple golang library for encrypting and decrypting data.

## Installation

To install the package, run:
```bash
go get -v github.com/deeper-x/encdec
```

## Usage

To encrypt data, run:
```go
import "github.com/deeper-x/encdec"

key := "<YOUR PASSWD>"
_, err := encdec.Encrypt("/path/to/resource", []byte(key))
if err != nil {
	panic(err)
}
```

To decrypt data, run:
```go
import "github.com/deeper-x/encdec"
            
key := "<YOUR PASSWD>"
_, err := encdec.Decrypt("/path/to/resource", []byte(key))
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
