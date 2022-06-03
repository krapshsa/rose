package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func GetItems() []*Item {
	var items = []*Item{
		{"apple", 5, 20},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Aged Brie", 10, 2},
		{"Aged Brie", 1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
		{"+5 Dexterity Vest", 10, 20},
		//{"Conjured Mana Cake", 20, 40},
	}
	return items
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}

}

const DayDecrease = 1
const DayExpired = 0
const QualUnit = 1
const QualMin = 0
const QualMax = 50

func UpdateQualityRev(items []*Item) {
	for _, item := range items {
		if item.Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		item.SellIn -= DayDecrease

		qualityVar := -QualUnit
		if item.SellIn < DayExpired {
			qualityVar *= 2
		}

		if item.Name == "Aged Brie" {
			qualityVar = QualUnit
			if item.SellIn < DayExpired {
				qualityVar += QualUnit
			}
		} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			qualityVar = QualUnit
			if item.SellIn < DayExpired {
				qualityVar = -item.Quality
			} else if item.SellIn < 5 {
				qualityVar += 2
			} else if item.SellIn < 10 {
				qualityVar += 1
			}
		} else if item.Name == "Conjured Mana Cake" {
			qualityVar *= 2
		}

		item.Quality += qualityVar

		if item.Quality < QualMin {
			item.Quality = QualMin
		} else if item.Quality > QualMax {
			item.Quality = QualMax
		}
	}
}
