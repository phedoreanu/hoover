# hoover
Hoover is a little Go room cleaning program.

###Installation
Hoover requires Go 1.5 or later.
```
go get -u github.com/phedoreanu/hoover
```

###Usage
Invoke `hoover` in the same directory as `input.txt`.

###Testing
To run the tests:
```
$ go test github.com/phedoreanu/hoover
ok      github.com/phedoreanu/hoover    0.005s
```
As always, if you are running the `go` tool from the package directory, you can omit the package path:
```
$ go test
PASS
ok      github.com/phedoreanu/hoover    0.005s
```
###Improvements
 * Add command-line flags for more verbose output and multiple input files
 * Docker + RESTful API
