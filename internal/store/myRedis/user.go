package myRedis

import "context"

func IsValidInviteCode(inviteCode string) (bool, error) {
	return Cli.SIsMember(context.Background(), "RegisterInviteCode", inviteCode).Result()
}

func AddInviteCode(inviteCode string) error {
	return Cli.SAdd(context.Background(), "RegisterInviteCode", inviteCode).Err()
}
