根據Guard Clauses 原則，對於明確的結果，先執行後略過，留下後面還要處理的

各個 item 都為獨立，以同層調列即可

將一些 magic number 定義成 const int

UpdateQuality 邏輯整理
- Sulfuras 不會變動
- 每一項物品，sellin - 1
- 不同類型物品，按照規則改變 quality
- 先初始化 quality 變化
- Aged Brie 與 Backstage passes 設定為 +1，根據條件更改
- 計算後的 quality，最後檢查有無超過上下限 

gildedrose/gildedrose.go => UpdateQualityRev 是將邏輯整理後的 refactoring

gildedrose/map.go => UpdateQualityMap 加入 function mapping 的概念，c style 的 refactroing
以 item name map 到對應的處理邏輯


gildedrose/interface.go => UpdateQualityTemplate 以物件導向的風格，參考 template pattern refactoring
go 並沒有繼承的概念，但是有類似 abstract 概念的 interface

ItemBase struct 含有一個 interface & Item
interface 規範需要實作 Update
每一個 item 需要轉換為 ItemBase，接著統一由 ItemBase 訂好的 update skelton 執行更新動作
skelton: 1. 各自 item update 2. post check

OO style 
