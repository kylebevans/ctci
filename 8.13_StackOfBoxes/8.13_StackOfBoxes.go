package main


type Box struct {
	w int
	h int
	d int
}

func SortBoxes(boxes []Box) []Box {
	lowmark := 0
	if len(boxes) <= 1 {
		return boxes
	}
	pivot := boxes[len(boxes) - 1].w
	for i := 0; i < len(boxes); i++ {
		if boxes[i].w <= pivot {
			boxes[lowmark], boxes[i] = boxes[i], boxes[lowmark]
			lowmark++
		}
	}
	left := SortBoxes(boxes[0:lowmark-1])
	right := SortBoxes(boxes[lowmark:])
	sortedBoxes := append(left, boxes[lowmark-1])
	sortedBoxes = append(sortedBoxes, right...)
	return sortedBoxes
}

	
func GetMaxStackHeight(boxes []Box) ([]Box, int) {
	var maxHeight int
	var maxStack []Box
	sortedBoxes := SortBoxes(boxes)
	for i := 0; i < len(sortedBoxes); i++ {
		stack := make([]Box, 0)
		var height int

		stack, height = CreateMaxStack(sortedBoxes[i+1:], sortedBoxes[i])
		if maxHeight < height {
			maxHeight = height
			maxStack = stack
		}
	}
	return maxStack, maxHeight		
}

func CreateMaxStack(boxes []Box, topBox Box) ([]Box, int) {
	maxStack := []Box{topBox}
	var maxHeight = topBox.h
	for i := 0; i < len(boxes); i++ {
		stack := make([]Box, 0)
		var height int
		if boxes[i].w > topBox.w && boxes[i].h > topBox.h && boxes[i].d > topBox.d {
			stack, height = CreateMaxStack(boxes[i+1:], boxes[i])
			if maxHeight < height + topBox.h {
				maxHeight = height + topBox.h
				tempStack := []Box{topBox}
				maxStack = append(tempStack, stack...)
			}
		}
	}
	return maxStack, maxHeight
}