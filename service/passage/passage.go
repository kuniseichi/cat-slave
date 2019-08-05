package passage

import (
	"cat-slave/model/passage"
	"cat-slave/service"
)

type PassageDto struct {

}

func PassageService(data PassageDto) error {
	return service.Transact(func () error {
		if _, err := passage.List(); err != nil {
			return err
		}
		if _, err := passage.Get(1); err != nil {
			return err
		}
		// 对传入的指针赋值
		data = 1
		return nil
	})
}