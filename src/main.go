package main

import (
	"fmt"
	"odev_four/lsm"
)

func main() {
	lsm := lsm.NewLSMTree(30, 30, "sstable_directory")

	fmt.Println("LSM Tree ve SSTable Kullanarak Veri Depolama Uygulaması")

	// Kullanıcı arayüzü ve işlevler
	for {
		fmt.Println("\n1. Ekle")
		fmt.Println("2. Getir")
		fmt.Println("3. Sil")
		fmt.Println("4. Çıkış")
		fmt.Print("Seçiminiz: ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Anahtarı girin: ")
			var key int
			if _, err := fmt.Scan(&key); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
				continue
			}
			fmt.Print("Değeri girin: ")
			var value string
			if _, err := fmt.Scan(&value); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir metin girin.")
				continue
			}
			lsm.Put(key, value)
			fmt.Println("Ekleme Başarılı.")
		case 2:
			fmt.Print("Anahtarı girin: ")
			var key int
			if _, err := fmt.Scan(&key); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
				continue
			}
			val, err := lsm.Get(key)
			if err == nil {
				fmt.Printf("Anahtar: %d, Değer: %s\n", key, val)
			} else {
				fmt.Printf("Anahtar bulunamadı: %d\n", key)
			}
		case 3:
			fmt.Print("Silmek istediğiniz anahtarı girin: ")
			var key int
			if _, err := fmt.Scan(&key); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
				continue
			}
			lsm.Delete(key)
			fmt.Printf("Anahtar %d silindi.\n", key)
		case 4:
			fmt.Println("Programdan çıkılıyor.")
			return
		default:
			fmt.Println("Geçersiz seçenek. Lütfen 1 ile 4 arasında bir sayı girin.")
		}
	}
}
