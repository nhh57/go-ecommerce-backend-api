package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

func GetUserKey(key string) string {
	return fmt.Sprintf("u:%s:otp", key)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// conver UUID to string , remove
	uuid := strings.ReplaceAll(newUUID.String(), "", "")
	return strconv.Itoa(userId) + "clitoken" + uuid
}
