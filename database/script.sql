CREATE TABLE go_crm_user
(
    usr_id               INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Account ID',
    usr_email            VARCHAR(255) NOT NULL COMMENT 'Email',
    usr_phone            VARCHAR(20)  NOT NULL COMMENT 'Phone',
    usr_username         VARCHAR(100) NOT NULL COMMENT 'Username',
    usr_password         VARCHAR(255) NOT NULL COMMENT 'Password',
    usr_created_at       INT          NOT NULL COMMENT 'Creation Time',
    usr_updated_at       INT          NOT NULL COMMENT 'Update Time',
    usr_create_ip_at     VARCHAR(45)  NOT NULL COMMENT 'Creation IP',
    usr_last_login_at    INT          NOT NULL COMMENT 'Last Login Time',
    usr_last_login_ip_at VARCHAR(45)  NOT NULL COMMENT 'Last Login IP',
    usr_login_times      INT          NOT NULL COMMENT 'Login Times',
    usr_status           TINYINT(1)   NOT NULL COMMENT 'Status 1:enable, 0:disable, -1:delete',
    KEY `inx_email` (`usr_email`),
    KEY `inx_phone` (`usr_phone`),
    KEY `inx_username` (`usr_username`)
) ENGINE = IGNORE
  DEFAULT CHARSET = utf8mb4 COMMENT = 'Account';
