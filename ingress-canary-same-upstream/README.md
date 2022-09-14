## 环境说明

共有三个后端服务
1. apple-app
2. peach-app
3. banana-app

fruit.bwangel.me 绑定在 apple-app 上
favorite.bwangel.me 绑定在 peach-app 上

canary 模式下

fruit.bwangel.me 正常情况下绑定在 banana-app 上
favorite.bwangel.me 绑定在 banana-app 上

## 要验证的问题

将此目录的代码和 [ingress-same-upstream](../ingress-same-upstream) 比较后可得，canary 模式下的 ingress 如果使用了同一个 upstream，那么 ingress 是不会出错的。
