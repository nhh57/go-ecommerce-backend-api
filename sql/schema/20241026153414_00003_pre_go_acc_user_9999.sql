-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_9999`
(
    `user_id`                BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
    `user_account`           VARCHAR(255)     NOT NULL COMMENT 'User account',
    `user_nickname`          VARCHAR(255)     NULL DEFAULT NULL COMMENT 'User nickname',
    `user_avatar`            VARCHAR(255)     NULL DEFAULT NULL COMMENT 'User avatar',
    `user_state`             TINYINT UNSIGNED NOT NULL COMMENT 'User state: 0-Locked ,  1-Activated ,  2-Not Activated',
    `user_mobile`            VARCHAR(20)      NULL DEFAULT NULL COMMENT 'Mobile phone number',
    `user_gender`            TINYINT UNSIGNED NULL DEFAULT NULL COMMENT 'User gender: 0-Secret ,  1-Male ,  2-Female',
    `user_birthday`          DATE             NULL DEFAULT NULL COMMENT 'User birthday',
    `user_email`             VARCHAR(255)     NULL DEFAULT NULL COMMENT 'User email address',
    `user_is_authentication` TINYINT UNSIGNED NOT NULL COMMENT 'Authentication status: 0-Not Authenticated ,  1-Pending ,  2-Authenticated ,  3-Failed',
    `created_at`             TIMESTAMP        NULL DEFAULT CURRENT_TIMESTAMP() COMMENT 'Record creation time',
    `updated_at`             TIMESTAMP        NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time',
    INDEX `idx_user_mobile` (`user_mobile`),
    INDEX `idx_user_email` (`user_email`),
    INDEX `idx_user_state` (`user_state`),
    INDEX `idx_user_is_authentication` (`user_is_authentication`)
) ENGINE = innoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = 'acc_user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_9999`;
-- +goose StatementEnd
