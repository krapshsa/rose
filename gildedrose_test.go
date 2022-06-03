package main

import (
	"fmt"
	"gildedrose/gildedrose"
	"testing"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		//t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func Test_Origin(t *testing.T) {
	/*var itemsa = []*gildedrose.Item{
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
	}*/
	itemsa := gildedrose.GetItems()
	itemsb := gildedrose.GetItems()

	for day := 0; day < 15; day++ {
		gildedrose.UpdateQuality(itemsa)
		gildedrose.UpdateQualityRev(itemsb)

		for i := 0; i < len(itemsa); i++ {
			fmt.Printf("day%d, %s left %d %d\t %d %d\n", day+1, itemsa[i].Name, itemsa[i].SellIn, itemsb[i].SellIn, itemsa[i].Quality, itemsb[i].Quality)
			if itemsa[i].Quality != itemsb[i].Quality {
				t.Errorf("After %d days, %d:%s quality a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].Quality, itemsb[i].Quality)
			}
		}
	}
}
