-- ----------------------------------------------------------------
-- Table structure for pre_go_acc_user_9999
-- ----------------------------------------------------------------

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

-- ----------------------------------------------------------------
-- Table structure for pre_go_acc_user_base_9999
-- ----------------------------------------------------------------
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

-- ----------------------------------------------------------------
-- Table structure for pre_go_acc_user_verify_9999
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_verify_9999`
(
    `verify_id`         INT          NOT NULL AUTO_INCREMENT,
    `verify_otp`        VARCHAR(6)   NOT NULL,
    `verify_key`        VARCHAR(255) NOT NULL,
    `verify_key_hash`   VARCHAR(255) NOT NULL,
    `verify_type`       INT               DEFAULT '1',
    `is_verified`       INT               DEFAULT '0',
    `is_deleted`        INT               DEFAULT '0',
    `verify_created_at` TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP(),
    `verify_updated_at` TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`verify_id`),
    UNIQUE KEY `unique_verify_key` (`verify_key`),
    INDEX `idx_verify_otp` (`verify_otp`)
) ENGINE = innoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = 'acc_user_verify';
SET FOREIGN_KEY_CHECKS = 1;