# cook-do

Cookpadレシピから指定の料理工程をTodoとして落とせるアプリです。
バックエンドはGoのAPI、フロントエンドはReactで動作します。
名前の由来は「Cookpad ToDoアプリ」を略して「Cook toDo!」です。

## ワイヤーフレーム、価値仮説シート、やることやらないことリスト

![cook-do](https://user-images.githubusercontent.com/22770924/44245071-fb5b2b80-a211-11e8-802e-aa38ec9fd466.JPG)

## Want

- Go1.9 ~
- mysql
- Docker


# Develop
```sh
$ make deps
```

## Test
### Go

```sh
$ make test

$ make integration-test

$ make api-test # Want run go server
```

### JavaScript

```sh
$ npm run test
```

## DB
default pass is `password`

### Initialize
```sh
$ make migrate/init

$ make migrate/init DBNAME=test-treasure  # Make test db
```

### Migrate
```sh
$ make migrate/dry  # Do not migrate

$ make migrate/up

$ make migrate/up ENV=test  # Test db migrate
```

### Status

```sh
$ make migrate/status
```

### MySQL Docker

```sh
$ make docker/start

or

$ make docker/stop
```
