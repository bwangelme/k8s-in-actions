## Bug 说明

共有三个后端服务
  1. apple-app
  2. peach-app
  3. banana-app
 
apple.bwangel.me 绑定在 apple-app 上
peach.bwangel.me 绑定在 peach-app 上
fruit.bwangel.me 正常情况下绑定在 apple-app 上
fruit.bwangel.me 在 canary 模式下，绑定到了 peach-app 上

## 期望的结果

```shell
# 访问的是 apple app
$ http fruit.bwangel.me
apple

# 访问的是 peach app
$ http fruit.bwangel.me "X-QAE-CANARY: always"
peach
```

## 实际的结果

```shell
# 访问的是 apple app
$ http fruit.bwangel.me
apple

# 访问的是 apple app
$ http fruit.bwangel.me "X-QAE-CANARY: always"
apple
```

## 纠正的手段

如果让 fruit.bwangel.me 在 canary 模式下，绑定到了 banana-app 上，那么就能正常工作

```shell
ø> http fruit.bwangel.me "X-QAE-CANARY: always"
banana

ø> http fruit.bwangel.me
apple
```

## 结论

```shell
NGINX Ingress controller
  Release:       v1.1.0
  Build:         cacbee86b6ccc45bde8ffc184521bed3022e7dee
  Repository:    https://github.com/kubernetes/ingress-nginx
  nginx version: nginx/1.19.9
```

在 1.1.0 版本的 ingress 中发现了这样的一行日志

2871:W0914 08:39:18.921819       7 controller.go:1517] alternative upstream xxxx in Ingress xxxxx is primary upstream in Other Ingress for location xxxxx/!

这说明 canary 模式的的 ingress, 它的 upstream 必须是独有的。