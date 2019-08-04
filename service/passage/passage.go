package passage

import (
	"cat-slave/model/passage"
	"cat-slave/service"
)

type PassageDto struct {

}

func PassageService() (PassageDto, error) {
	return service.Transact(func () (interface{}, error) {
		if _, err := passage.List(); err != nil {
			return err
		}
		if _, err := passage.Get(1); err != nil {
			return err
		}
		return nil
	})
}