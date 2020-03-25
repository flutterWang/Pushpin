package init

import "container/list"

func InitBasic() *List {
	var basicList = list.New()

	for i := 0; i < 1200; i++ {
		basicList.PushBack(1)
	}

	return basicList
}
