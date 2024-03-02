# Postfix+Golangを使用してGmailでSMTP送信する

## Docker Composeを使用した方法

### アプリケーション実行
```bash
$ docker-compose up --build
```

### Postfix内で確認(ログ確認)
```bash
$ docker exec -it postfix /bin/bash

$ cat /var/log/mail.log
```

### Postfix単独でメール送信
```bash
$ docker exec -it postfix /bin/bash

$ sendmail example@gmail.com
From:From@gmail.com
To:To@gmail.com
Subject:Hello

World

.
```

## Kindを使用した方法

### アプリケーションビルド
```bash
$ docker build -t app:latest --target prod ./app
```

### Postfixビルド
```bash
$ docker build -t postfix:latest .
```

### kind立ち上げ
```bash
$ kind create cluster --config=./manifest/kind.yaml
```


### クラスター内のノードにDockerイメージをロード
```bash
$ kind load docker-image app:latest
$ kind load docker-image postfix:latest
```
