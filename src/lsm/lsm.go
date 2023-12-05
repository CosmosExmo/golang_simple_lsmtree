package lsm

import (
	"fmt"
	"odev_four/utils"
	"sort"
)

type LSMTree struct {
	MemTable      map[int]string
	SSTables      []SSTable
	MaxMemTable   int
	MaxSSTable    int
	SSTablePrefix string
}

func NewLSMTree(maxMemTable, maxSSTable int, sstablePrefix string) *LSMTree {
	return &LSMTree{
		MemTable:      make(map[int]string),
		MaxMemTable:   maxMemTable,
		MaxSSTable:    maxSSTable,
		SSTablePrefix: sstablePrefix,
	}
}

func (lsm *LSMTree) Put(key int, value string) {
	lsm.MemTable[key] = value
	if len(lsm.MemTable) >= lsm.MaxMemTable {
		lsm.Compact()
	}
}

func (lsm *LSMTree) Get(key int) (string, error) {
	if val, ok := lsm.MemTable[key]; ok {
		return val, nil
	}
	for i := len(lsm.SSTables) - 1; i >= 0; i-- {
		if val, ok := lsm.SSTables[i].Data[key]; ok {
			return val, nil
		}
	}
	return "", fmt.Errorf("key not found: %d", key)
}

func (lsm *LSMTree) Delete(key int) {
	delete(lsm.MemTable, key)
}

func (lsm *LSMTree) Compact() {
	newSSTable := SSTable{
		Filename: lsm.getNextSSTableFilename(),
		Data:     make(map[int]string),
	}

	for key, value := range lsm.MemTable {
		newSSTable.Data[key] = value
	}

	lsm.SSTables = append(lsm.SSTables, newSSTable)

	sort.Slice(lsm.SSTables, func(i, j int) bool {
		return len(lsm.SSTables[i].Data) < len(lsm.SSTables[j].Data)
	})

	if len(lsm.SSTables) > lsm.MaxSSTable {
		lsm.SSTables = lsm.SSTables[1:]
	}

	if err := utils.WriteToFile(newSSTable.Filename, newSSTable.Data); err != nil {
		fmt.Println("Error writing SSTable to file:", err)
		return
	}

	lsm.MemTable = make(map[int]string)
}

func (lsm *LSMTree) getNextSSTableFilename() string {
	return fmt.Sprintf("%s/sstable_%d.txt", lsm.SSTablePrefix, len(lsm.SSTables))
}
