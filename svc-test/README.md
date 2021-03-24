NodePort Svc Test
===

## 说明

1. 创建2个 pod，只绑定在1个节点上 (k8s-node2)
2. 创建一个指向这两个 Pod 的 Nodeport service
3. 看是不是只有两个节点创建了 port

## 结论

k8s-node1 和 k8s-node3 都监听了 30080 这个端口
但是由于设置了 externalTrafficPolicy，nc 访问不了这两个节点的 30080 端口

关于 externalTrafficPolcy 的文章: https://www.cnblogs.com/zisefeizhu/p/13262239.html
