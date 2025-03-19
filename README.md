See:

- [wit-bindgen-go](./main.go) directive in main.go
- [Custom Option](./cm/custom.go) with json serialization

```
cm-override ❯ curl -i http://127.0.0.1:8000
HTTP/1.1 200 OK
vary: origin, access-control-request-method, access-control-request-headers
access-control-allow-origin: *
access-control-expose-headers: *
connection: close
transfer-encoding: chunked
date: Wed, 19 Mar 2025 20:43:39 GMT

marshal
{"string-value":"hello","int-value":42,"bool-value":true,"list-value":["val1","val2"]}

unmarshal
{"string-value":"hello","int-value":42,"bool-value":true,"list-value":["val1","val2"]}

marshal with nulls
{"string-value":null,"int-value":null,"bool-value":null,"list-value":null}

unmarshal with nulls
{"string-value":null,"int-value":null,"bool-value":null,"list-value":null}
```
