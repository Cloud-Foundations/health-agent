package sysfs

func ReadUint64(filename string) (uint64, error) {
	return readUint64(filename)
}

func ReadBool(filename string) (bool, error) {
	return readBool(filename)
}

func ReadString(filename string) (string, error) {
	return readString(filename)
}
