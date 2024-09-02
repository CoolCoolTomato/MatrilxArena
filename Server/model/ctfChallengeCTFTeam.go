package model

import (
    "gorm.io/gorm"
    "time"
)

type CTFChallengeCTFTeam struct {
    CTFChallengeID uint
    CTFTeamID uint
    SolvedAt   time.Time
}

func (u *CTFChallengeCTFTeam) BeforeCreate(tx *gorm.DB) (err error) {
    u.SolvedAt = time.Now()
    return
}

func (u *CTFChallengeCTFTeam) BeforeUpdate(tx *gorm.DB) (err error) {
    u.SolvedAt = time.Now()
    return
}
