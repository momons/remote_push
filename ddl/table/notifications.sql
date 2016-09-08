-- アプリ通知マスタ
CREATE TABLE notifications (
    -- ID
    id bigserial PRIMARY KEY,
    -- 通知コード
    notification_code VARCHAR(64) NOT NULL,
    -- 通知名
    notification_name VARCHAR(256) NOT NULL,
    -- certファイル名(iOS)
    cert_file_name VARCHAR(256) NOT NULL,
    -- keyファイル名(iOS)
    key_file_name VARCHAR(256) NOT NULL,
    -- APIキー(Android)
    api_key VARCHAR(256) NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
