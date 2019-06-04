package box

import (
	"math/rand"
	"time"
)

type ItemType string

const (
	// BOXITEM 宝箱
	BOXITEM ItemType = "box"
)

// Item : 箱子中的奖励项
type Item struct {
	T      ItemType // 奖励类型
	Des    int64    // 奖励描述（例如：金币的数量，宝箱的id）
	Weight int64    // 奖励权重，（-1）表示必开到
}

// Box : 宝箱
type Box struct {
	Items []*Item
}

// Open : 开箱
func (b *Box) Open() []*Item {
	// 首先将必须发的奖励拿出来
	var givelist []*Item
	var randlist []*Item
	for _, item := range b.Items {
		if item.Weight == -1 {
			givelist = append(givelist, item)
		} else if item.Weight == 0 {
			// dnt give this item
		} else {
			if item.T == BOXITEM {
				// 通过 item.Des 拿到宝箱的具体内容，然后进行开箱，将其放到givelist中
				b := Box{
					Items: []*Item{
						&Item{T: "Cloth", Des: 101, Weight: 10},
						&Item{T: "Hat", Des: 1, Weight: 110},
					},
				}
				givelist = append(givelist, b.Open()...)
			} else {
				randlist = append(randlist, item)
			}
		}
	}

	type a struct {
		i      *Item
		offset int64
		count  int64
	}
	aLists := make([]*a, 0, len(randlist))
	var sumCount int64
	for _, item := range randlist {
		aLists = append(aLists, &a{
			i:      item,
			offset: sumCount,
			count:  item.Weight,
		})
		sumCount += item.Weight
	}
	if sumCount == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano() + 561612)
	awardIndex := rand.Int63n(sumCount)
	for _, u := range aLists {
		if u.offset+u.count > awardIndex {
			givelist = append(givelist, u.i)
			break
		}
	}
	return givelist
}
