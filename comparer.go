package comparer

import "fmt"

func AreEqual(list1, list2 []string) (bool, int) {
	big := list1
	small := list2
	if len(list2) > len(list1) {
		big = list2
		small = list1
	}

	fmt.Printf("big: %#v\n", big)
	fmt.Printf("small: %#v\n", small)

	var deltaBig, deltaSmall int
	for i := 0; i < len(big); {
		fmt.Printf("\ni: %d\n", i)

		indexBig := i + deltaBig
		indexSmall := i + deltaSmall
		fmt.Printf("indexBig: %d\n", indexBig)
		fmt.Printf("indexSmall: %d\n", indexSmall)

		if len(small) == indexSmall || len(big) == indexBig {
			return true, 0
		}

		if len(small)-1 < indexSmall {
			fmt.Printf("Big (%d) has more elements than small (%d)\n", len(big), len(small))
			return false, i
		}

		stringBig := big[indexBig]
		stringSmall := small[indexSmall]

		if stringBig == "" {
			fmt.Printf("Skip empty string in big[%d]\n", indexBig)
			deltaBig++
			continue
		}

		if stringSmall == "" {
			fmt.Printf("Skip empty string in small[%d]\n", indexSmall)
			deltaSmall++
			continue
		}

		if stringBig != stringSmall {
			return false, i
		}

		fmt.Println("Equals. Next")
		i++
	}

	return true, 0
}
