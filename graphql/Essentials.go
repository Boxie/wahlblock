package graphql

//TODO Check if length = int
func calculatePaging(offset int, first int, length int) (int, int){

	if offset < 0 || first <= 0 {
		return 0,0
	}

	if offset < length {

		if offset + first < length{
			return offset, offset + first - 1
		}

		return offset, length - 1

	}

	return 0,0
}
