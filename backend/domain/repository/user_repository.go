package repository

import (
	"context"

	"github.com/MISW/Portal/backend/domain"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

// UserRepository - User関連のDB操作
type UserRepository interface {
	// Insert - 新規サークル員の追加
	Insert(ctx context.Context, user *domain.User) (int, error)

	// GetByID - IDで検索
	GetByID(ctx context.Context, id int) (*domain.User, error)

	// GetByID - Slack IDで検索
	GetBySlackID(ctx context.Context, slackID string) (*domain.User, error)

	// GetByEmail - Emailで検索
	GetByEmail(ctx context.Context, email string) (*domain.User, error)

	// List - 全ユーザを取得
	List(ctx context.Context) ([]*domain.User, error)

	// ListByID - ユーザIDが一致する全てのユーザを取得
	ListByID(ctx context.Context, ids []int) ([]*domain.User, error)

	// Update - ユーザのプロフィールを更新する(idで識別)
	Update(ctx context.Context, user *domain.User) error

	// VerifyEmail - メールアドレスを認証済みにする
	VerifyEmail(ctx context.Context, id int, email string) error
}
