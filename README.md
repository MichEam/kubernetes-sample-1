## Kubernetes Ingressによるリクエスト振り分けサンプル

```
nginx.$(minikube ip).nip.io/echo/*  --> echo-app-service:8080
nginx.$(minikube ip).nip.io/hello/* --> hello-app-service:8080
nginx.$(minikube ip).nip.io/        --> nginx-service:80
```

## 参考
- [Minikube で快適に Ingress を利用する - Qiita](https://qiita.com/superbrothers/items/13d8ce012ef23e22cb74)
- [Minikubeで快適にIngressを利用したいが、dnsmasqの設定は面倒なので省略したい - Qiita](https://qiita.com/nobusue/items/4817c19c0279f070c24b)

