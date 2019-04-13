package usecase

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/usecase/interfaces"
)

type InputInteractor struct {
	InputRepository interfaces.InputRepository
	Logger          interfaces.Logger
}

func (i *InputInteractor) Get1() (*domain.Input, error) {
	//i.Logger.Log("get data!")
	//return i.InputRepository.Get1()

	tmp, err := i.InputRepository.Get1()

	v := reflect.Indirect(reflect.ValueOf(tmp.CreatedAt))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		// フィールド名
		println("Field: " + t.Field(i).Name)
	}

	fmt.Println(tmp.CreatedAt)
	fmt.Println(time.Date(2019, time.April, 11, 22, 22, 22, 43259, time.Local))

	return tmp, err
}

/*
func (i *InputInteractor) Get10(input *domain.Input) (int, error) {
	i.Logger.Log("get data!")
	return i.InputRepository.Get10(input)
}
*/
