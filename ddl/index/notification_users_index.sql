-- 通知ユーザデータテーブル ユニークインデックス
CREATE UNIQUE INDEX notification_users_index ON notification_users (
    notification_code,
    notification_token
);
