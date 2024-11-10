# Quiz App


## Notes

### General
I need to make less stuff public, using testing has gotten me into the habit of making everything public so I can test it.


### Creating My First Web App
[link](https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html)

Starts by Making a simple web server
- Hello World

Moves on to template library.

### net/http package
[link](https://pkg.go.dev/net/http)

This is a really well documented package, go here if you need more info. 

The http package provides HTTP client and server implementations.

```
resp, err := http.Get(path)
resp, err := http.Post(path, contentType, body)
resp, err := http.PostForm(url, data)
```

caller must close the body when done with it.

Clients and Transports: Controled by this library

Servers
----
built in
```
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```
custom
```
s :+ &http.Server{
    Addr:           ":8080",
    Handler:        myHandler,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20m,
}

log.Fatal(s.ListenAndServe())

```


### Research Handlers and Servemuxes

Handlers
- Handle application logic and writing response headers and bodies

Servermux
- Router
- Mapping between the predefined URL paths and corresponding handlers. There is generally one ServerMux for your application containing all routes.
- Stands for: HTTP Request Multiplexer


### Templates

[syntax doc](https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax)

[Better Docs](https://gohugo.io/templates/introduction/)


