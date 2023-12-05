package lsm

type SSTable struct {
	Filename string
	Data     map[int]string
}
