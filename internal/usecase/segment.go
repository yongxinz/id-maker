package usecase

import (
	"id-maker/internal/entity"
	"id-maker/pkg/snowflake"
)

// SegmentUseCase -.
type SegmentUseCase struct {
	repo      SegmentRepo
	alloc     *Alloc
	snowFlake *snowflake.Worker
}

// New -.
func New(r SegmentRepo) *SegmentUseCase {
	var err error

	s := &SegmentUseCase{}
	s.repo = r
	if s.alloc, err = s.NewAllocId(); err != nil {
		panic(err)
	}
	if s.snowFlake, err = s.NewAllocSnowFlakeId(); err != nil {
		panic(err)
	}
	return s
}

func (uc *SegmentUseCase) GetId(tag string) (id int64, err error) {
	uc.alloc.Mu.Lock()
	defer uc.alloc.Mu.Unlock()

	val, ok := uc.alloc.BizTagMap[tag]
	if !ok {
		// tag 不存在则创建
		if err = uc.CreateTag(&entity.Segments{
			BizTag: tag,
			MaxId:  1,
			Step:   10000,
		}); err != nil {
			return 0, err
		}
		val, _ = uc.alloc.BizTagMap[tag]
	}
	return val.GetId(uc)
}

// SnowFlakeGetId -.
func (uc *SegmentUseCase) SnowFlakeGetId() int64 {
	return uc.snowFlake.GetId()
}

// CreateTag -.
func (uc *SegmentUseCase) CreateTag(e *entity.Segments) (err error) {
	if err = uc.repo.Add(e); err != nil {
		return
	}
	b := &BizAlloc{
		BazTag:  e.BizTag,
		GetDb:   false,
		IdArray: make([]*IdArray, 0),
	}
	b.IdArray = append(b.IdArray, &IdArray{
		Cur:   1,
		Start: 0,
		End:   e.Step,
	})
	uc.alloc.BizTagMap[e.BizTag] = b
	return
}
