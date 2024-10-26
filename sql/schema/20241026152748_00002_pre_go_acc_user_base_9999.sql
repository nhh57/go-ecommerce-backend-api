-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_base_9999`
(
    `user_id`          INT          NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_account`     VARCHAR(255) NOT NULL,
    `user_password`    VARCHAR(255) NOT NULL,
    `user_salt`        VARCHAR(255) NOT NULL,
    `user_login_time`  TIMESTAMP    NULL DEFAULT NULL,
    `user_logout_time` TIMESTAMP    NULL DEFAULT NULL,
    `user_login_ip`    VARCHAR(45)  NULL DEFAULT NULL,
    `user_created_at`  TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP(),
    `user_updated_at`  TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `unique_user_account` (`user_account`)
) ENGINE = innoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = 'acc_user_base';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd
