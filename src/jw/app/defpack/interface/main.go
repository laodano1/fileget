package main

import "fmt"

type Player struct {
	gold int
}

type MetaItem struct {
	id int
	v0 int
	mainType int
}

type IItem interface {
	Count() int
}

type IEquipment interface {
	IItem
	WearPart() int
}

type IUsableItem interface {
	IItem
	use(p *Player) bool
}

type Item struct {
	mi *MetaItem
	count int
}

func (i *Item) Count() int {
	return i.count
}

type Equipment struct {
	Item
}

func (i *Equipment) WearPart() int {
	return i.mi.v0
}

type UsableItem struct {
	*Item
}

func (i *UsableItem) use(p *Player) bool {
	p.gold++
	i.count--
	return true
}

type Bag struct {
	items []IItem
}

func (bag *Bag) get(idx int) IItem {
	return bag.items[idx]
}

func NewItem(mi *MetaItem, count int) IItem {
	switch mi.mainType {
	case 0:
		return &Equipment{Item{mi:mi,count:count}}
	case 1:
		return &UsableItem{&Item{mi:mi,count:count}}
	}
	return nil
}

func init() {
	p := Player{}
	b := Bag{}
	b.items = make([]IItem, 10)
	mi := &MetaItem{}
	mi.mainType = 1
	b.items[0] = NewItem(mi, 1)

	// in GameSession::use_item_handler function
	i, ok := b.get(0).(IUsableItem)
	if !ok {
		return
	}

	if i.Count() >= 1 {
		i.use(&p)
	}

	fmt.Println(i,p.gold)
}
