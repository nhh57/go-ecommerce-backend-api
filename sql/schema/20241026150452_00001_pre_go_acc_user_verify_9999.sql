-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_verify_9999`
(
    `verify_id`         INT          NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `verify_otp`        VARCHAR(6)   NOT NULL, -- ID of the OTP record
    `verify_key`        VARCHAR(255) NOT NULL, -- OTP code
    `verify_key_hash`   VARCHAR(255) NOT NULL,
    `verify_type`       INT               DEFAULT '1',
    `is_verified`       INT               DEFAULT '0',
    `is_deleted`        INT               DEFAULT '0',
    `verify_created_at` TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP(),
    `verify_updated_at` TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_verify_otp` (`verify_otp`),
    UNIQUE KEY `unique_verify_key` (`verify_key`)
) ENGINE = innoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = 'acc_user_verify';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
-- +goose StatementEnd
