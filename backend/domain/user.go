package domain

import (
	"regexp"
	"strings"
	"time"

	"github.com/MISW/Portal/backend/internal/rest"
	"golang.org/x/xerrors"
)

// SexType - 性別
type SexType string

const (
	// Men - 男性
	Men SexType = "men"
	// Women - 女性
	Women SexType = "women"
	// Other - その他
	Other SexType = "other" // not used
)

// RoleType - サークル員の種別
type RoleType string

const (
	// Admin - 管理者(会員資格あり)
	Admin RoleType = "admin"

	// Member - 正式なメンバー(会員資格あり)
	Member RoleType = "member"

	// Retired - 引退済みのメンバー(会員資格あり)
	Retired RoleType = "retired"

	// NotMember - 未払い状態のメンバー(会員資格なし)
	NotMember RoleType = "not_member"
)

// University - 所属大学
type University struct {
	Name       string `json:"name" yaml:"name"`
	Department string `json:"department" yaml:"department"`
	Subject    string `json:"subject" yaml:"subject"`
}

// User - サークル員の情報
type User struct {
	ID                   int         `json:"id" yaml:"id"`
	Email                string      `json:"email" yaml:"email"`
	Generation           int         `json:"generation" yaml:"generation"`
	Name                 string      `json:"name" yaml:"name"`
	Kana                 string      `json:"kana" yaml:"kana"`
	Handle               string      `json:"handle" yaml:"handle"`
	Sex                  SexType     `json:"sex" yaml:"sex"`
	University           *University `json:"university" yaml:"university"`
	StudentID            string      `json:"student_id" yaml:"student_id"`
	EmergencyPhoneNumber string      `json:"emergency_phone_number" yaml:"emergency_phone_number"`
	OtherCircles         string      `json:"other_circles" yaml:"other_circles"`
	Workshops            []string    `json:"workshops" yaml:"workshops"`
	Squads               []string    `json:"squads" yaml:"squads"`
	Role                 RoleType    `json:"role" yaml:"role"`

	// 外部サービス
	SlackID string `json:"slack_id" yaml:"slack_id"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`
}

var (
	emailValidator = regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)

	invalidWordsForSquads = []string{"\n", "\r"}
)

func (user *User) Validate() error {
	for i := range user.Squads {
		for j := range invalidWordsForSquads {
			if strings.Contains(user.Squads[i], invalidWordsForSquads[j]) {
				return rest.NewBadRequest("班の名前に使えない文字を含んでいます")
			}
		}
	}

	if user.Sex != Men && user.Sex != Women {
		return rest.NewBadRequest("性別の値が不正です")
	}
	if !emailValidator.MatchString(user.Email) {
		return rest.NewBadRequest("メールアドレスの形式が不正です")
	}

	return nil
}

var (
	// ErrEmailConflicts - emailが既に登録されている
	ErrEmailConflicts = xerrors.New("email conflicts")

	// ErrSlackIDConflicts - Slack IDが既に登録されている
	ErrSlackIDConflicts = xerrors.New("slack id conflicts")

	// ErrNoUser - Userが存在しない
	ErrNoUser = xerrors.New("no such user")
)
