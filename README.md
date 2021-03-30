# Go URL

## Introduction

GoURL is written as IE Spring 2021 course project to simulate the curl behaviour in Go.

## Examples

```sh
> ./gourl https://httpbin.org/get

INFO[0000] sending request into https://httpbin.org/get
INFO[0000] Method is: GET
INFO[0000] Response status: 200 OK
INFO[0000] headers are:
INFO[0000] Date = [Tue, 30 Mar 2021 15:40:08 GMT]
INFO[0000] Content-Type = [application/json]
INFO[0000] Content-Length = [273]
INFO[0000] Server = [gunicorn/19.9.0]
INFO[0000] Access-Control-Allow-Origin = [*]
INFO[0000] Access-Control-Allow-Credentials = [true]
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-60634658-7e2795bc70bfe6e7286c4f0d"
  },
  "origin": "94.139.160.188",
  "url": "https://httpbin.org/get"
}
```

```sh
> ./gourl https://httpbin.org/post -X POST -j '{ "hello": "world" }'

INFO[0000] sending request into https://httpbin.org/post
INFO[0001] Method is: POST
INFO[0001] Response status: 200 OK
INFO[0001] headers are:
INFO[0001] Content-Length = [453]
INFO[0001] Server = [gunicorn/19.9.0]
INFO[0001] Access-Control-Allow-Origin = [*]
INFO[0001] Access-Control-Allow-Credentials = [true]
INFO[0001] Date = [Tue, 30 Mar 2021 15:43:37 GMT]
INFO[0001] Content-Type = [application/json]
{
  "args": {},
  "data": "{ \"hello\": \"world\" }",
  "files": {},
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "20",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-60634729-41ac79314b2e7c9775a86f39"
  },
  "json": {
    "hello": "world"
  },
  "origin": "94.139.160.188",
  "url": "https://httpbin.org/post"
}
```
