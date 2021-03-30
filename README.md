# Go URL

## Introduction

GoURL is written as IE Spring 2021 course project to simulate the curl behaviour in Go.

## Examples

1. GET Request
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

2. POST Request with JSON body and `Content-Type` automatically is set to `application/json`.
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

3. POST Request with Form body and `Content-Type` automatically is set to `application/x-www-form-urlencoded`.
```sh
> ./gourl https://httpbin.org/post -X POST -d 'hello=world'

INFO[0000] sending request into https://httpbin.org/post
INFO[0001] Method is: POST
INFO[0001] Response status: 200 OK
INFO[0001] headers are:
INFO[0001] Date = [Tue, 30 Mar 2021 15:44:52 GMT]
INFO[0001] Content-Type = [application/json]
INFO[0001] Content-Length = [448]
INFO[0001] Server = [gunicorn/19.9.0]
INFO[0001] Access-Control-Allow-Origin = [*]
INFO[0001] Access-Control-Allow-Credentials = [true]
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "hello": "world"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "11",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-60634774-4bc210632c9ea0bd2c0ccce0"
  },
  "json": null,
  "origin": "94.139.160.188",
  "url": "https://httpbin.org/post"
}
```
4. GET Request with Query Strings

```sh
> ./gourl https://httpbin.org/get -X GET -Q 'hello:world' -Q 'salam:donya'

INFO[0000] sending request into https://httpbin.org/get?hello=world&salam=donya
INFO[0001] Method is: GET
INFO[0001] Response status: 200 OK
INFO[0001] headers are:
INFO[0001] Date = [Tue, 30 Mar 2021 15:48:31 GMT]
INFO[0001] Content-Type = [application/json]
INFO[0001] Content-Length = [344]
INFO[0001] Server = [gunicorn/19.9.0]
INFO[0001] Access-Control-Allow-Origin = [*]
INFO[0001] Access-Control-Allow-Credentials = [true]
{
  "args": {
    "hello": "world",
    "salam": "donya"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-6063484f-15f9b9f903601d420c6fb621"
  },
  "origin": "94.139.160.188",
  "url": "https://httpbin.org/get?hello=world&salam=donya"
}
```
