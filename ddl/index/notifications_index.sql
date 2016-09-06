-- 通知マスタテーブル ユニークインデックス
CREATE UNIQUE INDEX notifications_index ON notifications (
    notification_code
);
