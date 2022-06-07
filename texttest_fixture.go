package main

import (
	"bufio"
	"fmt"
	"gildedrose/gildedrose"
	"os"
	"strconv"
)

func main() {
	var result string = "OMGHAI!\n"

	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 31
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		result += fmt.Sprintf("-------- day %d --------\n", day)
		result += fmt.Sprintln("name, sellIn, quality")
		for i := 0; i < len(items); i++ {
			result += fmt.Sprintf("%s, %d, %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		}
		result += fmt.Sprintln("")
		gildedrose.UpdateQuality(items)
	}

	f, _ := os.Create("./gildedrose_test_actual.txt")
	w := bufio.NewWriter(f)
	_, _ = w.WriteString(result)
	_ = w.Flush()
}
