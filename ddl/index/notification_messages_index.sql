--  通知メッセージデータテーブル ユニークインデックス
CREATE UNIQUE INDEX notification_messages_index ON notification_messages (
    message_code
);
