package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	// Mydate Date
	Date	// embedding
	/*
		has a ( 포함 ) -> embedding, 맴버랑 다른 게 뭐 임
		is a ( 상속 )
		interface: 규격
	*/
}

func (e *Event) SetTitle(title string) error {

	if utf8.RuneCountInString(title) > 6 {	// 글자의 개수를 새는 것 ( 크기 X )
		return errors.New("이벤트는 6글자 이내로 작성해야 합니다.")
	}

	e.title = title
	return nil
}

func (e *Event) GetTitle() string{
	return e.title
}
