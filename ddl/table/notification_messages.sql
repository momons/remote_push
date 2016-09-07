-- 通知メッセージデータ
CREATE TABLE notification_messages (
    -- ID
    id bigserial PRIMARY KEY,
    -- 通知メッセージコード
    message_code VARCHAR(64) NOT NULL,
    -- 通知コード
    notification_code VARCHAR(64) NOT NULL,
    -- 送信日時
    send_at CHAR(14) NOT NULL,
    -- プラットフォーム
    platform VARCHAR(64) NOT NULL,
    -- 送信メッセージ
    message TEXT NOT NULL,
    -- 送信条件 (JSON)
    send_condition TEXT NOT NULL,
    -- 送信済み
    is_send CHAR(1) NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
