package src

type Play struct {
	playerFields   []int
	computerFields []int
	usedFields     []int
}

func NewPlay() *Play {
	return &Play{[]int{}, []int{}, []int{}}
}

func (p *Play) MakePlayerMove(fieldNum int) {
	p.playerFields = append(p.playerFields, fieldNum)
	p.usedFields = append(p.usedFields, fieldNum)
}

func (p *Play) MakeComputerMove(fieldNum int) {
	p.computerFields = append(p.playerFields, fieldNum)
	p.usedFields = append(p.usedFields, fieldNum)
}
