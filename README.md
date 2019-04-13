# About

自動でエアコンの On/Off するやつ

* 一定間隔ラズパイで温度を取って DB に突っ込む & 赤外線センサを操作するREST APIの作成
* DB のデータを得る REST API の作成 & 赤外線センサ用 REST API を叩く REST API の作成 (Web Backend)
    * ここ
* ブラウザから温湿度のグラフ、On/Off のボタン表示 (Web Frontend)
* 温湿度情報を出力する Slack Bot

# Run

* run immediately
```
go run main.go
```

* build
```
go build -o $(basename $PWD) main.go
```

* test
```
go test -v ./... -cover
```


# Ref

* http://nakawatch.hatenablog.com/entry/2018/07/11/181453
* https://qiita.com/tenntenn/items/24fc34ec0c31f6474e6d

---
