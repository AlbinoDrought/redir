# Redir

Run:

```sh
docker build -t ghcr.io/albinodrought/redir .
docker run --rm -it \
  -p 80:8080 \
  -e "REDIR_SERVER=nginx/1.10.1" \
  -e "REDIR_FALLBACK=172.22.0.4" \
  ghcr.io/albinodrought/redir
```

Check:

```sh
curl -v -h "Host: " http://localhost:80
```

```
> GET / HTTP/1.1
> Host: 
> User-Agent: curl/8.4.0
> Accept: */*
> 
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: https://172.22.0.4/
< Server: nginx/1.10.1
< Date: Tue, 05 Dec 2023 05:43:55 GMT
< Content-Length: 54
< 
<a href="https://172.22.0.4/">Moved Permanently</a>.
```