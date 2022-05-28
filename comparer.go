package comparer

func AreEqual(list1, list2 []string) (bool, int) {
	bigList := list1
	smallList := list2
	if len(list2) > len(list1) {
		bigList = list2
		smallList = list1
	}

	bigMap := mapList(bigList)
	smallMap := mapList(smallList)

	for i := 0; i < len(bigMap); i++ {
		lineBig := bigMap[i]
		lineSmall, ok := smallMap[i]
		if !ok || lineBig != lineSmall {
			return false, lineBig.index
		}
	}

	return true, 0
}

type line struct {
	index   int
	content string
}

func mapList(list []string) map[int]line {
	lines := map[int]line{}
	mapIndex := 0
	for i := 0; i < len(list); i++ {
		item := list[i]
		if item == "" {
			continue
		}

		lines[mapIndex] = line{
			index:   i,
			content: item,
		}
		mapIndex++
	}
	return lines
}
