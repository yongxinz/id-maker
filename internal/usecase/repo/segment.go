package repo

import (
	"fmt"
	"id-maker/internal/entity"
	"id-maker/pkg/mysql"
	"time"
)

// SegmentRepo -.
type SegmentRepo struct {
	*mysql.Mysql
}

// New -.
func New(mysql *mysql.Mysql) *SegmentRepo {
	return &SegmentRepo{mysql}
}

// GetList -.
func (r *SegmentRepo) GetList() ([]entity.Segments, error) {
	var s []entity.Segments

	if err := r.Engine.Find(&s); err != nil {
		return s, fmt.Errorf("SegmentRepo - GetList - Find: %w", err)
	}

	return s, nil
}

// Add -.
func (r *SegmentRepo) Add(s *entity.Segments) error {
	var (
		exist bool
		err   error
	)

	if exist, err = r.Engine.Where("biz_tag = ?", s.BizTag).Exist(&entity.Segments{}); err != nil {
		return fmt.Errorf("SegmentRepo - Add - Exist: %w", err)
	}
	if exist {
		return fmt.Errorf("Tag Already Exist")
	}
	if _, err = r.Engine.Insert(s); err != nil {
		return fmt.Errorf("SegmentRepo - Add - Insert: %w", err)
	}

	return nil
}

// GetNextId -.
func (r *SegmentRepo) GetNextId(tag string) (*entity.Segments, error) {
	var (
		err error
		id  = &entity.Segments{}
		tx  = r.Engine.Prepare()
	)

	if _, err = tx.Exec(
		"update segments set max_id=max_id+step, update_time = ? where biz_tag = ?", time.Now(), tag); err != nil {
		_ = tx.Rollback()
		return id, fmt.Errorf("SegmentRepo - GetNextId - Exec: %w", err)
	}
	if _, err = tx.Where("biz_tag = ?", tag).Get(id); err != nil {
		_ = tx.Rollback()
		return id, fmt.Errorf("SegmentRepo - GetNextId - Get: %w", err)
	}
	err = tx.Commit()

	return id, nil
}
