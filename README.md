# Postfix+Golangを使用してGmailでSMTP送信する


## アプリケーション実行
```bash
$ docker-compose up --build
```

## Postfix内で確認(ログ確認)
```bash
$ docker exec -it postfix /bin/bash

$ cat /var/log/mail.log
```

## Postfix単独でメール送信
```bash
$ docker exec -it postfix /bin/bash

$ sendmail example@gmail.com
From:From@gmail.com
To:To@gmail.com
Subject:Hello

World

.
```
