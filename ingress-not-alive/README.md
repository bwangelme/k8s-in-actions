## 实验

banana-app pod 是 0，请求 canary 时，走到了默认的 main ingress 中

```shell
http fruit.bwangel.me "X-QAE-CANARY: always"
apple
```

ingress-controller 日志中有报错

```shell
2022/09/14 10:27:09 [error] 1483#1483: *2268768 [lua] balancer.lua:203: route_to_alternative_balancer(): no alternative balancer for backend: qae-banana-service-8000, client: 192.168.56.1, server: fruit.bwangel.me, request: "GET / HTTP/1.0", host: "fruit.bwangel.me"
```

将 banana-app 的 pod 启动起来之后，canary 就能正常工作了

```shell
ø> http fruit.bwangel.me "X-QAE-CANARY: always"
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 7
Content-Type: text/plain; charset=utf-8
Date: Wed, 14 Sep 2022 10:29:35 GMT
Request-Id: 4d8efcb87656c707c98b13fc3c8a50fc
Server: nginx/1.18.0 (Ubuntu)
X-App-Name: http-echo
X-App-Version: 0.2.3
X-QAE: NORMAL

banana

```