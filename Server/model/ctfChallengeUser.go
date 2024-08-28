package model

import (
    "gorm.io/gorm"
    "time"
)

type CTFChallengeUser struct {
    UserID         uint
    CTFChallengeID uint
    SolvedAt   time.Time
}

func (u *CTFChallengeUser) BeforeCreate(tx *gorm.DB) (err error) {
    u.SolvedAt = time.Now()
    return
}

func (u *CTFChallengeUser) BeforeUpdate(tx *gorm.DB) (err error) {
    u.SolvedAt = time.Now()
    return
}
