package db

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// GetUser returns user playlistId and count of user's added songs
func (s *Service) GetUser(ctx context.Context, userId int64) (string, int, error) {
	user, err := s.db.Get(ctx, strconv.FormatInt(userId, 10)).Result()

	if err != nil {
		return "", 0, err
	}

	userData := strings.Split(user, ";")
	playlistId := userData[0]
	tracksCount, err := strconv.Atoi(userData[1])

	if err != nil {
		return "", 0, err
	}

	return playlistId, tracksCount, nil
}

func (s *Service) UpdateUser(ctx context.Context, userId int64, playlistId string, tracksCount int) error {
	data := fmt.Sprintf("%s;%d", playlistId, tracksCount)
	return s.db.Set(ctx, strconv.FormatInt(userId, 10), data, 0).Err()
}
