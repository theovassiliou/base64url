# bae64url

base64url, implemented in go (golang) is a small library supporting base64url encoding and decoding (collectively called coding). In addition, 
it is a command-line tool similar to BSD's `base64` tool.

WARNING: THIS SOFTWARE CAN'T BE ERROR-FREE, SO USE IT AT YOUR OWN RISK. I HAVE DONE MY BEST TO MAKE SURE THAT THE TOOLS BEHAVE AS EXPECTED. BUT AGAIN ... USE IT AT YOUR OWN RISK. I AM NOT GIVING ANY KIND OF WARRANTY, NEITHER EXPLICITLY NOR IMPLICITLY.

## Installation binaries

You can download the binaries directly from the [releases](https://github.com/theovassiliou/base64url/releases) section.  Unzip/untar the downloaded archive and copy the files to a location of your choice, e.g. `/usr/local/bin/` on *NIX or MacOS. If you install only the binaries, make sure that they are accessible from the command line. Ideally, they are accessible via `$PATH` or `%PATH%`, respectively.

## Installation From Source

base64url requires golang version 1.13 or newer, the Makefile requires GNU make.

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

There is no particular requirement beyond the fact that you should have a working go installation.

[Install Go](https://golang.org/doc/install) >=1.13

### Installing

Download base64url source by running

```shell
go get -u github.com/theovassiliou/bae64url
```

This gets you your copy of base64url installed under
`$GOPATH/src/github.com/theovassiliou/base64url`

Run `make` from the source directory by running

```shell
cd $GOPATH/src/github.com/theovassiliou/base64url
make all
```

to compile and build the executable

* base64url - [README](cmd/base64url/README.md)

and run

```shell
make go-install
```

to install a copy of the executable into `$GOPATH/bin`


## Running the tests

We are using two different make targets for running tests.

```shell
make test
go test -short `go list`
ok      github.com/theovassiliou/base64url 0.3s
```

executes all short package tests, while

```shell
make test-all
go vet $(go list ./...)
go test ./...
?    github.com/theovassiliou/hc2-tools/cmd/expandRequire [no test files]
?    github.com/theovassiliou/hc2-tools/cmd/hc2DownloadScene [no test files]
?    github.com/theovassiliou/hc2-tools/cmd/hc2SceneInteract [no test files]
?    github.com/theovassiliou/hc2-tools/cmd/hc2UploadScene [no test files]
ok   github.com/theovassiliou/hc2-tools/pkg 0.036s
```

executes in addition `go vet`on the package. Before committing to the code base please use `make test-all` to ensure that all tests pass.

### Break down into end to end tests

After creating your configuration call `hc2DownloadScene` without `-u -p` parameters.

```shell
hc2DownloadScene -t

Successful connected to ...
  Name         : Hal
  Serial       : HC2-033533
  IP           : 192.10.66.55
  Version      : 4.560
  ZWaveVersion : 3.67

and logged in as:
  User:         specialuser@mydomain.com
  Type:         superuser
```

to test whether command can execute correctly.

## Deployment

After running

```shell
make install

go build -ldflags " -X main.commit=99f909d -X main.branch=master" ./cmd/hc2UploadScene
go build -ldflags " -X main.commit=99f909d -X main.branch=master" ./cmd/hc2DownloadScene
go build -ldflags " -X main.commit=99f909d -X main.branch=master" ./cmd/hc2SceneInteract
mkdir -p /usr/local/bin/
cp hc2UploadScene /usr/local/bin/
cp hc2DownloadScene /usr/local/bin/
cp hc2SceneInteract /usr/local/bin/
```

you can find your executables in `/usr/local/bin`. Make sure `/usr/local/bin/` is in your path.

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/theovassiliou/hc2-tools/tags).

## Authors

* **Theo Vassiliou** - *Initial work* - [Theo Vassiliou](https://github.com/theovassiliou)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

Thanks to all the people out there that produce amazing open-source software, which supported the creation of this piece of software. In particular I wasn't only able to use libraries etc. But also, to learn and understand golang better. In particular I wanted to thank

* [Jaime Pillora](https://github.com/jpillora) for [jpillora/opts](https://github.com/jpillora/opts). Nice piece of work!
* [InfluxData Team](https://github.com/influxdata) for [influxdata/telegraf](https://github.com/influxdata/telegraf). Here I learned a lot for Makefile writing and release building in particular.
* Inspiration and motivation to develop this tool I got from the [ZeroBrane](https://studio.zerobrane.com/) Lua Development Environment.
* [PurpleBooth](https://gist.github.com/PurpleBooth) for the well motivated [README-template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)

***

## History

This project has been developed as I was seeking for a way to upload scenes from my favorite development Lua development environment to the Fibaro HC2 system. Finally I came up with the idea to upload a scene whenever I do a `git commit`. For this I needed a cmd line tool can be integrated as `commithook` into the git repository.

With this I could solve two problems at a single time.

1. Enforcing a version control system, e.g. git
2. Automatically uploading the modified script

After implementing a first version, new ideas emerged, so for example retrieving debug messages where implemented.
