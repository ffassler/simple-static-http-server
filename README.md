# Simple static http server

The aim of this project is to offer a simple [http](https://fr.wikipedia.org/wiki/Hypertext_Transfer_Protocol) server with :
- just one feature : serve static files
- simple to be configured : command line arguments
- light : small binary size
- very easy to deploy : copy the binary and that it ! No package, no node or python are required.

To realize this aim, the server is written in [go](https://golang.org/) with no additional modules.

Because of that, for example, under a Linux OS the required dependies are only :

```sh
$ ldd simple-static-http-server 
	linux-vdso.so.1 (0x00007ffd3d8e6000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f1c79ab6000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f1c796c5000)
	/lib64/ld-linux-x86-64.so.2 (0x00007f1c79cd5000)
```

The size is also very small. For example, under Ubuntu 18.04 :

```
$ ll simple-static-http-server 
-rwxrwxr-x 1 ubuntu01 ubuntu01 6.5M Feb 12 13:53 simple-static-http-server
```

## How to build from sources

### Prerequisites

- [go binary](https://golang.org/doc/install)

### Build

```sh
go build -o simple-static-http-server
```

## Usage

All configurations are defined by command line arguments.
These arguments is printed by the flag -h

```sh

$ ./simple-static-http-server -h
Usage of ./simple-static-http-server:
  -d string
    	Directory which contains static files to be served (default "/home/ubuntu01/workspace-go/simple-static-http-server")
  -p int
    	Port (default 8080)
```
The default directory to be served is the current directory and the default is 8080.

To change it, use the flag -d and -p.

For example, to serve the directory /staic on port 9090.

```
$ ./simple-static-http-server -d /static -p 9090
```

In case of error, logs are printed.

```
$ ./simple-static-http-server -d /static -p 80  
2019/02/12 14:17:04 Cannot start the http server : listen tcp :80: bind: permission denied
```

## Deployment

Just copy the binary !

## Docker

A docker image is available on dockerhub [squareitservices/static-http-server](https://hub.docker.com/r/squareitservices/static-http-server)

The docker image is based on alpine with just the only one required dependecy *libc6-compat*

### Build image from dockefile

The [dockerfile](Dockerfile) is a multi stage build which build the binary from sources and then create the docker image

```sh
docker build -t static-http-server:tag .
```

The docker image is also light :

```sh
$ docker images | grep squareitservices/static-http-server
squareitservices/static-http-server                              latest              7b91cb0d4b15        3 days ago          12.3MB
```
