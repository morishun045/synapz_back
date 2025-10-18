-- データベースセットアップスクリプト
-- MySQLに接続後、以下のコマンドを実行してください

-- データベースの作成
CREATE DATABASE IF NOT EXISTS synapz_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 専用ユーザーの作成
CREATE USER IF NOT EXISTS 'synapz_user'@'localhost' IDENTIFIED BY 'synapz_password';
GRANT ALL PRIVILEGES ON synapz_db.* TO 'synapz_user'@'localhost';
FLUSH PRIVILEGES;

-- 使用するデータベースを選択
USE synapz_db;

-- データベースが正常に作成されたか確認
SHOW TABLES;