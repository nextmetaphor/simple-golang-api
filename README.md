# Simple `golang` API
This repository contains an extremely simple API written in `golang` which simply returns a JSON document detailing the various attributes of the request. Also included are additional instructions for building a Docker image based on [Alpine Linux](https://alpinelinux.org/) to run this API.

## Getting Started

### Prerequisites
* Local [golang](https://golang.org/) installation; see [https://nextmetaphor.io/2016/12/09/getting-started-with-golang-on-macos/](https://nextmetaphor.io/2016/12/09/getting-started-with-golang-on-macos/) for details on how to install on macOS
* Local [Docker](https://www.docker.com/) installation

### Install

#### Building the Code
To build the code, it is as simple as running the commands below.<br><br>
*Note that the first two `export` commands are only required if you plan to run the API within the Alpine Linux Docker container; if this is not the case then these can be omitted.*
```
export GOOS=linux
export GOARCH=amd64

go build
```

#### Building The Docker Image
Following the completion of the steps above, the Docker image can be built as follows, changing the tag name if required. 
```
docker build . -t nextmetaphor/simple-golang-api:latest
```

## Deployment

### Running The API From Docker
To create a container based on the image previously built, simply run the command below. This will map container port `8080` to `28080` on the local machine.
```
docker run -p 28080:8080 -d nextmetaphor/simple-golang-api
```

## Validation

### Testing the API
Assuming that the API is being run from Docker, simply use `cURL` to hit the root of the API.
```
curl http://localhost:28080/
```

This should return a document similar to that shown below.
```
{"Method":"GET","URL":{"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/","RawPath":"","ForceQuery":false,"RawQuery":"","Fragment":""},"Proto":"HTTP/1.1","ProtoMajor":1,"ProtoMinor":1,"Header":{"Accept":["*/*"],"User-Agent":["curl/7.51.0"]},"Body":{},"ContentLength":0,"TransferEncoding":null,"Close":false,"Host":"localhost:28080","Form":null,"PostForm":null,"MultipartForm":null,"Trailer":null,"RemoteAddr":"172.17.0.1:44534","RequestURI":"/","TLS":null,"Response":null}
```

## Licence ##
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.