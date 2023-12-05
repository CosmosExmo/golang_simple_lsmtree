package lsm

import (
	"fmt"
	"odev_four/utils"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type LSMTree struct {
	MemTable      map[int]string
	SSTables      []SSTable
	MaxMemTable   int
	MaxSSTable    int
	SSTablePrefix string
}

func NewLSMTree(maxMemTable, maxSSTable int, sstablePrefix string) *LSMTree {
	lsmTree := &LSMTree{
		MemTable:      make(map[int]string),
		MaxMemTable:   maxMemTable,
		MaxSSTable:    maxSSTable,
		SSTablePrefix: sstablePrefix,
	}

	// SSTable dosyalarını tara ve listeye ekle.
	lsmTree.loadSSTables()

	return lsmTree
}

func (lsm *LSMTree) loadSSTables() {
	filepath.Walk(lsm.SSTablePrefix, func(path string, info os.FileInfo, err error) error {
			if err != nil {
					fmt.Println("Error accessing path:", path)
					return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".txt") {
					// Dosya ismini SSTable listesine ekle.
					lsm.SSTables = append(lsm.SSTables, SSTable{Filename: path})
			}

			return nil
	})
}


func (lsm *LSMTree) Put(key int, value string) {
	lsm.MemTable[key] = value
	if len(lsm.MemTable) >= lsm.MaxMemTable {
		lsm.Compact()
	}
}

func (lsm *LSMTree) Get(key int) (string, error) {
	// Önce MemTable'da ara.
	if val, ok := lsm.MemTable[key]; ok {
		return val, nil
	}
	
	// SSTable'larda ara.
	for i := len(lsm.SSTables) - 1; i >= 0; i-- {
		sstable := &lsm.SSTables[i]
		
		// Eğer SSTable bellekte yoksa, dosyadan oku.
		if sstable.Data == nil {
			data, err := utils.ReadFromFile(sstable.Filename)
			if err != nil {
				return "", fmt.Errorf("error reading SSTable from file: %s, error: %v", sstable.Filename, err)
			}
			sstable.Data = data
		}

		// Bellekteki SSTable'da anahtarı ara.
		if val, ok := sstable.Data[key]; ok {
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
