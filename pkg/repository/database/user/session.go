package user

import (
	"context"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"gorm.io/gorm"
	"time"
)

func (r *Repository) CreateSession(ctx context.Context, session entities.Session) error {
	repoSession := FromDomainToSession(session)
	return r.DB.WithContext(ctx).Create(&repoSession).Error
}

func (r *Repository) GetSessionByToken(ctx context.Context, token string) (entities.Session, error) {
	var repoSession SessionRepo
	err := r.DB.WithContext(ctx).Where("token = ?", token).First(&repoSession).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.Session{}, nil
		}
		return entities.Session{}, err
	}
	return repoSession.ToDomain(), nil
}

func (r *Repository) DeleteSession(ctx context.Context, token string) error {
	return r.DB.WithContext(ctx).Where("token = ?", token).Delete(&SessionRepo{}).Error
}

func (r *Repository) DeleteExpiredSessions(ctx context.Context) error {
	return r.DB.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&SessionRepo{}).Error
}
