-- 通知ユーザデータ
CREATE TABLE notification_users (
    -- ID
    id bigserial PRIMARY KEY,
    -- 通知コード
    notification_code VARCHAR(64) NOT NULL,
    -- 通知トークン
    notification_token VARCHAR(256) NOT NULL,
    -- プラットフォーム
    platform VARCHAR(64) NOT NULL,
    -- カスタムパラメータ
    custom_params TEXT NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
