package fate

import (
	"github.com/globalsign/mgo/bson"
	"github.com/godcong/chronos"
	"github.com/godcong/fate/mongo"
	"github.com/godcong/yi"
	"strconv"
)

//Name 姓名
type Name struct {
	FirstName   []*Character //名姓
	LastName    []*Character
	born        *chronos.Calendar
	baGua       *yi.Yi //周易八卦
	zodiac      *Zodiac
	zodiacPoint int
}

func (n Name) String() string {
	var s string
	for _, l := range n.LastName {
		s += l.Ch
	}
	for _, f := range n.FirstName {
		s += f.Ch
	}
	return s
}

func (n Name) WuXing() string {
	var s string
	for _, l := range n.LastName {
		s += l.WuXing
	}
	for _, f := range n.FirstName {
		s += f.WuXing
	}
	return s
}

func createName(impl *fateImpl, f1 *Character, f2 *Character) *Name {
	lastSize := len(impl.lastChar)
	last := make([]*Character, lastSize, lastSize)
	copy(last, impl.lastChar)

	shang := getStroke(last[0])
	if lastSize > 1 {
		shang += getStroke(last[1])
	}

	ff1 := *f1
	ff2 := *f2
	first := []*Character{&ff1, &ff2}

	xia := getStroke(first[0]) + getStroke(first[1])

	return &Name{
		FirstName: first,
		LastName:  impl.lastChar,
		baGua:     yi.NumberQiGua(xia, shang, shang+xia),
	}
}

//MakeName input the lastname to make a name
func FilterName(names []*Name) *Name {
	//todo:make name generate list

	return &Name{}
}

//BaGua
func (n *Name) BaGua() *yi.Yi {
	return n.baGua
}

func nameCharacter(s string) *mongo.Character {
	c := mongo.Character{}
	err := mongo.C("character").Find(bson.M{
		"character": s,
	}).One(&c)
	log.Infof("%+v", c)
	if err != nil {
		return nil
	}
	return &c
}

//CountStroke 统计笔画
func CountStroke(chars ...*mongo.Character) int {
	i := 0
	if chars == nil {
		return i
	}
	for k := range chars {
		t, _ := strconv.Atoi(chars[k].KangxiStrokes)
		i += t
	}
	return i
}
