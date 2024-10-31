-- name: AddUserHaveUserId :execresult
INSERT INTO `pre_go_acc_user_9999` (user_id, user_account, user_nickname, user_avatar, user_state, user_mobile,
user_gender, user_birthday, user_email, user_is_authentication)
VALUES  (?,?,?,?,?,?,?,?,?,?);