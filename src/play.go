package src

import "fyne.io/fyne/v2/widget"

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
	id   int
	view *widget.Button
}

func NewField(id int, view *widget.Button) *Field {
	return &Field{id: id, view: view}
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
	p.computerFields = append(p.playerFields, fieldNum)
	p.usedFields = append(p.usedFields, fieldNum)

	p.markField(fieldNum, "X")
}

func (p *Play) markField(fieldNum int, symbol string) {
	field := p.fields[fieldNum]

	if "" == field.view.Text {
		field.view.SetText(symbol)
		field.view.Disable()
	}
}
