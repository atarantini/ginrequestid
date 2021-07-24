# Gingonic X-Request-Id middleware

Basic middleware to add and expose an [UUID4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29>) for each request based on the code posted on [Dan Sosedoff development and experiments](https://sosedoff.com/2014/12/21/gin-middleware.html) website.

It will set the variable ``RequestId`` or order to use it from within the application for logging or propagation.

If ``X-Request-Id`` header is sent in the request, it will set that value as ``RequestId``.


## Installation

Use ```go get``` to fetch the package:

```bash
$ go get github.com/atarantini/ginrequestid
```


## Usage

Import it and use it in your router:

```golang
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/atarantini/ginrequestid"
    ...
)

func main() {
    // Initialize router
    r := gin.Default()

    // Load middleware
    r.Use(ginrequestid.RequestId())

    // Your routes
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    // Start server
    r.Run()
}
```

Check if it's working:

```bash
$ curl --verbose localhost:8080/ping
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /ping HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=utf-8
< X-Request-Id: 3528d29d-f5c6-4758-ae05-eee5a56c27ea
< Date: Wed, 25 Jan 2017 12:42:30 GMT
< Content-Length: 4
<
* Connection #0 to host localhost left intact
```

Use it for logging::

```golang
r.GET("/ping", func(c *gin.Context) {
    fmt.Println(fmt.Sprintf("[request-id:%s][event:ping]", c.MustGet("RequestId")))
    c.String(200, "pong")
})
```

## Changelog

``0.0.3`` - 2021-07-24

* Updated to a new version of UUID library, please this don't change AGAIN :|

``0.0.2`` - 2018-03-06

* Fix for change in signature in the main dependecy of the package. In order to use the last version of `github.com/satori/go.uuid` update to `0.0.2` version.


``0.0.1`` - 2017-02-07

* Initial release, basic middleware. It works!


## Author

Andres Tarantini (atarantini@gmail.com)


## License

GNU GPL v3. See LICENSE file.
