package src

import (
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

	p.MakeComputerMove(p.getComputerField())
}

func (p *Play) MakeComputerMove(fieldNum int) {
	time.Sleep(1 * time.Second)
	p.computerFields = append(p.playerFields, fieldNum)
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

func (p *Play) getComputerField() int {
	chosenField := p.fields[rand.Intn(9)]

	if false == chosenField.isAvailable {
		return p.getComputerField()
	}

	return chosenField.id
}
