Ingress Canary Example
===

这是一个 Ingress 的 Cannary 用法的而例子。

客户端如果传了 X-QAE-CANARY header，那么请求就会走到 qae-stage 中，否则请求会走到 qae 中。