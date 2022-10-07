package enum

import (
	"errors"
	"strings"
)

type GameType struct {
	gameType string
}

func (gt *GameType) String() string {
	return gt.gameType
}

var (
	Unknown = GameType{"unknown"}
	Wheel   = GameType{"lucky_wheel"}
)

func (gt *GameType) FromString(s string) (GameType, error) {
	s = strings.ToLower(s)
	switch s {
	case Wheel.gameType:
		return Wheel, nil
	default:
		return Unknown, errors.New("unknown game type")
	}
}
