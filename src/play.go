package src

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

type Play struct {
	playerFields   []int
	computerFields []int
	usedFields     []int
	fields         [9]*Field
}

func (p *Play) SetFields(fields [9]*Field) {
	p.fields = fields
}

type Field struct {
	id          int
	view        *widget.Button
	isAvailable bool
}

func (f Field) IsAvailable() bool {
	return f.isAvailable
}

func NewField(id int, view *widget.Button) *Field {
	return &Field{id: id, view: view, isAvailable: true}
}

func NewPlay() *Play {
	return &Play{[]int{}, []int{}, []int{}, [9]*Field{}}
}

func (p *Play) MakePlayerMove(fieldNum int) {
	p.playerFields = append(p.playerFields, fieldNum)
	p.usedFields = append(p.usedFields, fieldNum)

	p.markField(fieldNum, "0")
}

func (p *Play) MakeComputerMove(fieldNum int) {
	time.Sleep(1 * time.Second)
	p.computerFields = append(p.computerFields, fieldNum)
	p.usedFields = append(p.usedFields, fieldNum)

	p.markField(fieldNum, "X")
}

func (p *Play) markField(fieldNum int, symbol string) {
	field := p.fields[fieldNum]

	if "" == field.view.Text {
		field.view.SetText(symbol)
		field.isAvailable = false
		field.view.Disable()
	}
}

func (p *Play) GetComputerField() int {
	rand.Seed(time.Now().UnixNano())
	chosenField := p.fields[rand.Intn(9)]

	if false == chosenField.isAvailable {
		return p.GetComputerField()
	}

	return chosenField.id
}

func (p *Play) DidPlayerWin(fieldNum int) bool {
	if p.isMinimumAmountOfFieldsUsed(p.playerFields) {
		return false
	}

	return p.isCorrectPatternMade(fieldNum, p.playerFields)
}

func (p *Play) DidComputerWin(fieldNum int) bool {
	if p.isMinimumAmountOfFieldsUsed(p.computerFields) {
		return false
	}

	return p.isCorrectPatternMade(fieldNum, p.computerFields)
}

func (p *Play) isCorrectPatternMade(fieldNum int, fields []int) bool {
	switch fieldNum {
	case 0:
		return p.containsSliceExpectedValues(fields, 1, 2) || p.containsSliceExpectedValues(
			fields,
			3,
			6,
		) || p.containsSliceExpectedValues(fields, 4, 8)
	case 1:
		return p.containsSliceExpectedValues(fields, 0, 2) || p.containsSliceExpectedValues(fields, 4, 7)
	case 2:
		return p.containsSliceExpectedValues(fields, 0, 1) || p.containsSliceExpectedValues(
			fields,
			5,
			8,
		) || p.containsSliceExpectedValues(fields, 4, 6)
	case 3:
		return p.containsSliceExpectedValues(fields, 0, 6) || p.containsSliceExpectedValues(fields, 4, 5)
	case 4:
		return p.containsSliceExpectedValues(fields, 0, 8) || p.containsSliceExpectedValues(
			fields,
			1,
			7,
		) || p.containsSliceExpectedValues(fields, 2, 6) || p.containsSliceExpectedValues(fields, 3, 5)
	case 5:
		return p.containsSliceExpectedValues(fields, 2, 8) || p.containsSliceExpectedValues(fields, 3, 4)
	case 6:
		return p.containsSliceExpectedValues(fields, 0, 3) || p.containsSliceExpectedValues(
			fields,
			7,
			8,
		) || p.containsSliceExpectedValues(fields, 4, 2)
	case 7:
		return p.containsSliceExpectedValues(fields, 6, 8) || p.containsSliceExpectedValues(fields, 1, 4)
	case 8:
		return p.containsSliceExpectedValues(fields, 0, 4) || p.containsSliceExpectedValues(
			fields,
			6,
			7,
		) || p.containsSliceExpectedValues(fields, 2, 5)
	}

	panic(fmt.Sprintf("Unexpected value: %d", fieldNum))
}

func (p *Play) containsSliceExpectedValues(fields []int, values ...int) bool {
	var contains []bool

	for _, field := range fields {
		for _, value := range values {
			if field == value {
				contains = append(contains, true)
			}
		}
	}

	return 2 == len(contains)
}

func (p *Play) isMinimumAmountOfFieldsUsed(fields []int) bool {
	return 3 > len(fields)
}
