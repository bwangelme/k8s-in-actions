## Ingress Canary on Same Namespace

apple-ing 和 banana-ing 这两个 ingress 在同一个 namespace 中，有一个设置了 `nginx.ingress.kubernetes.io/canary: "true"`，此时还是能正常工作的

```shell
ø> http fruit.bwangel.me "X-QAE-CANARY: always"
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 7
Content-Type: text/plain; charset=utf-8
Date: Wed, 14 Sep 2022 08:31:36 GMT
Request-Id: 0f9cd9b418ddb39eba72de8fd5144ba0
Server: nginx/1.18.0 (Ubuntu)
X-App-Name: http-echo
X-App-Version: 0.2.3
X-QAE: NORMAL

banana


ø> http fruit.bwangel.me
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 6
Content-Type: text/plain; charset=utf-8
Date: Wed, 14 Sep 2022 08:31:41 GMT
Request-Id: 43b0b54de76f02078aedd394d5f676ba
Server: nginx/1.18.0 (Ubuntu)
X-App-Name: http-echo
X-App-Version: 0.2.3
X-QAE: NORMAL

apple

```