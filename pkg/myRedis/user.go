package myRedis

import "context"

func IsInviteCode(inviteCode string) (bool, error) {
	return Cli.SIsMember(context.Background(), "RegisterInviteCode", inviteCode).Result()
}

func SetInviteCode(inviteCode string) error {
	return Cli.SAdd(context.Background(), "RegisterInviteCode", inviteCode).Err()
}
