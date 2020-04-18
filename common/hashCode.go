package common

//hash算法
func GetHashCode(array []byte) uint64 {
	hash := uint64(5381)
	keyLen := len(array) - 1
	for ; keyLen >= 8; keyLen -= 8 {
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
	}
	switch keyLen {
	case 7:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 6:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 5:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 4:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 3:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 2:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 1:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		fallthrough
	case 0:
		hash = ((hash << 5) + hash) + uint64(array[keyLen])
		keyLen--
		break
	}
	return hash
}
