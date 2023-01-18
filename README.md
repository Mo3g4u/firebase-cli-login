# firebase-cli-login

## 使い方

### login

```sh
$ cd login
$ go run main.go [FirebaseApiKey] [tenantId] [email] [password]
```

ログインに成功するとidToken(jwt)が表示されます

### refresh

```sh
$ cd refresh
$ go run main.go [FirebaseApiKey] [refreshToken]
```

ログインに成功するとidToken(jwt)が表示されます