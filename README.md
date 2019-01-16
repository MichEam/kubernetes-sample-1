# kubernetes-sample-1

Kubernetes Ingressをつかったデプロイメントのサンプル。

nginx-ingress-controllerによって、
URLパスごとに別のServiceへリクエストを振り分けている。

```
Rules:
  Host                         Path  Backends
  ----                         ----  --------
  nginx.192.168.99.100.nip.io
                               /echo    echo-app-service:8080 (<none>)
                               /hello   hello-app-service:8080 (<none>)
                               /        nginx-service:80 (<none>)                               
```

また、振り分け時にコンテキストを削除している。

```
http://nginx.192.168.99.100.nip.io/echo/some/path

↓↓↓↓↓↓↓

http://echo-app:8080/some/path
```

## 参考
- [Minikube で快適に Ingress を利用する - Qiita](https://qiita.com/superbrothers/items/13d8ce012ef23e22cb74)
- [Minikubeで快適にIngressを利用したいが、dnsmasqの設定は面倒なので省略したい - Qiita](https://qiita.com/nobusue/items/4817c19c0279f070c24b)
