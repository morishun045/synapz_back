# synapz

## concept

## about this repository

## contributors

## architecture

## セットアップガイド

### データベースセットアップ

#### 1. MySQLサーバーの起動
```bash
# MySQLサーバーを起動
sudo service mysql start
# または
sudo systemctl start mysql
```

#### 2. データベースとユーザーの作成
```bash
# rootユーザーでMySQLに接続
mysql -u root -p

# setup_db.sqlスクリプトを実行
mysql> source setup_db.sql;
# または
mysql -u root -p < setup_db.sql
```

#### 3. アプリケーションの実行
```bash
go run main.go
```

### 設定内容

- **データベース名**: `synapz_db`
- **ユーザー名**: `synapz_user` 
- **パスワード**: `synapz_password`
- **ホスト**: `localhost:3306`

### 環境変数での設定

本番環境では`.env`ファイルを作成して環境変数を設定することを推奨します：

```bash
cp .env.example .env
# .envファイルを編集して適切な値を設定
```

### トラブルシューティング

#### MySQL接続エラーの場合
1. MySQLサーバーが起動しているか確認
2. `setup_db.sql`が正常に実行されたか確認
3. ユーザーとデータベースが作成されているか確認：
   ```sql
   SHOW DATABASES;
   SELECT user, host FROM mysql.user WHERE user = 'synapz_user';
   ```
