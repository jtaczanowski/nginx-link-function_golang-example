# nginx-link-function_golang-example
Example of nginx plugin written in Golang and loaded by nginx-link-function

# usage

## configure nginx
Add nginx-link-function https://github.com/Taymindis/nginx-link-function module to your nginx.

## Inside repository directory compile example golang code as C shared library
```go build -o main.so -buildmode=c-shared .```

## load your plugin in nginx.conf

```
# nginx.conf
server {
  listen 8989;
  aio threads;
  ...
  ngx_link_func_lib "/path/to/your-plugin/main.so"; # sharing data memory with server 1 if the path are same with server 1
  ...
  ...
  location = /test {
      ngx_link_func_call "return_custom_http_headers_and_response" 
  }
}
```

