# A reverse proxy for GAE/SE

This reverse proxy is implemented by using [appengine urlfetch](https://cloud.google.com/appengine/docs/standard/go/issue-requests).
The [appengine log](https://cloud.google.com/appengine/articles/logging) package is also available by default.

## How to use

The implementation is fundamental, in fact you will customize `Director` and `ModifyResponse`.
That would be something like:

```go
package your_package

import (
  "log"
  "net/url"
  "net/http"

  revproxy "github.com/namusyaka/reverseproxy_appengine"
)

func init() {
  backendUrl := "https://example.appspot.com"
  u, err := url.Parse(backendUrl)
  if err != nil {
    log.Fatal(err)
  }
  p := revproxy.NewSingleHostReverseProxy(u)

  // Customize Director
  p.Director = func(r *http.Request) {
    // do something 
  }
  p.ModifyResponse = func(r *http.Response) error {
    // do something 
  }
  http.Handle("/", p)
}
```

## License

the MIT License
