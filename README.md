# ASAKATSUSNS

**朝活に特化したSNSアプリです**

  **URL:** https://front.asakatsusns.com

# アプリの概要
基本的にはTwitterのような、投稿・コメント・いいね・フォロー機能のあるSNSですがそれに追加して
以下のような特徴のあるアプリケーションにしました。
  - 目標起床時間を設定して、早起き達成日数を記録することができる
  - 投稿にタグ付けしタグごとに朝活の共有ができる。

# 開発した背景
私自身1日を有意義に過ごすには、朝の動きがとても大事であると最近強く感じており、それは科学的にも証明されています。
意識的に朝活を行うことでユーザーのQOLを上げたいと思い　、朝活に特化したSNSを作成しようと考えました。

# 使用技術
- フロントエンド
  - Vue.js 3.2.26
  - CSS
  - Bootstrap5.0
  - vue-fontawesome3.0

- バックエンド(API)
  - Golang 1.15.2
  - Air (ホットリロード)
  - Gin
  - Gorm
  - firebase authentication

- インフラ
  - CircleCi 2.1
  - Docker 20.10.12/ docker-compose 2.2.3
  - mysql 5.7.36
  - AWS (EC2, ALB, ACM, S3, RDS, S3, ECS, ECR, CloudFront, Route53, CloudWatch, VPC, IAM, RDS)

- その他使用ツール
  - Visual Studio Code

# 機能一覧
- ユーザー登録関連
  - アカウント新規登録、プロフィール編集機能
  - ログイン、ログアウト機能(JWT・firebase authentication)

- ユーザーの早起き達成日数のランキング機能

- 無限スクロール機能(v3-infinite-loading)

- ページネーション機能
  - コメント一覧、フォロー中/フォロワー一覧など

- コメント機能
  - コメント一覧機能（ページネーション）

- タグ機能 (vue3-tags-input)
  - タグ毎の投稿一覧機能

- いいね機能 (Vue.js / ajax)
  - いいねした投稿一覧機能

- フォロー機能
  - フォロー中/フォロワー一覧（ページネーション）

- アイコン画像アップロード機能
  - データベースにパスを保存し、S3に画像を保存。

- レスポンシブデザイン

# AWS構成図
<img width="691" alt="インフラ構成図" src="https://user-images.githubusercontent.com/66294061/154385303-6845f0ea-2034-45b4-9ec4-6ba5c88101e5.png">

# DB設計
<img width="1062" alt="ER図" src="https://user-images.githubusercontent.com/66294061/154876168-d7bd2502-e947-42fa-83bc-7a1675814a7e.png">
