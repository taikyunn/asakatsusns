# ASAKATSUSNS

**朝活に特化したSNSアプリです**
 **URL:** https://front.asakatsusns.com

# アプリの概要

# 開発した背景

# イメージ

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
 - Firebase Authentication

- インフラ
 - CircleCi
 - Docker 20.10.12/ docker-compose 2.2.3
 - mysql 5.7.36
 - AWS (EC2, ALB, ACM, S3, RDS, S3, ECS, ECR, CloudFront, Route53, CloudWatch, VPC, IAM, RDS)

- その他使用ツール
 - Visual Studio Code

# AWS構成図

# 機能一覧
- 一般ユーザー登録関連
  - アカウント新規登録、プロフィール編集機能
  - ログイン、ログアウト機能(JWT)

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

- レスポンシブデザイン
# DB設計

# 各テーブルについて



