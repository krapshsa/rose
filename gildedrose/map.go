package gildedrose

var fnMap = map[string]func(item *Item){
	"Sulfuras, Hand of Ragnaros":                fnSulf,
	"Aged Brie":                                 fnBrie,
	"Backstage passes to a TAFKAL80ETC concert": fnPass,
	"Conjured Mana Cake":                        fnConj,
}

func fnSulf(item *Item) {
	// do nothing
}

func fnBrie(item *Item) {
	item.SellIn -= DayDecrease
	item.Quality += QualUnit
	if item.SellIn < DayExpired {
		item.Quality += QualUnit
	}
	postCheck(item)
}

func fnPass(item *Item) {
	item.SellIn -= DayDecrease

	if item.SellIn < DayExpired {
		item.Quality = QualMin
		return
	}

	item.Quality += QualUnit
	if item.SellIn < 5 {
		item.Quality += 2
	} else if item.SellIn < 10 {
		item.Quality += 1
	}
	postCheck(item)
}

func fnConj(item *Item) {
	item.SellIn -= DayDecrease
	item.Quality -= 2

	if item.SellIn < DayExpired {
		item.Quality -= 2
	}
	postCheck(item)
}

func fnNormal(item *Item) {
	item.SellIn -= DayDecrease
	item.Quality -= QualUnit
	if item.SellIn < DayExpired {
		item.Quality -= QualUnit
	}
	postCheck(item)
}

func postCheck(item *Item) {
	if item.Quality < QualMin {
		item.Quality = QualMin
	} else if item.Quality > QualMax {
		item.Quality = QualMax
	}
}

func UpdateQualityMap(items []*Item) {
	for _, item := range items {
		if fn, ok := fnMap[item.Name]; ok {
			fn(item)
		} else {
			fnNormal(item)
		}
	}
}
