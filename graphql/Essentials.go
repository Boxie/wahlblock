package graphql

//TODO Check if length = int
func calculatePaging(offset int, first int, length int) (int, int){
	if offset > length - 1{
		return length, length
	}

	if offset + first > length - 1 {
		return offset, length
	}

	return offset, first + offset
}
