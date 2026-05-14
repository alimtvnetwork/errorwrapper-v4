package refs

func AllIndividualItemsCount(first *Collection, collections ...*Collection) int {
	length := 0

	if first != nil {
		length += first.Length()
	}

	for _, collection := range collections {
		if collection == nil {
			continue
		}

		length += collection.Length()
	}

	return length
}
