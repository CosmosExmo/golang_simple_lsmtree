package main

import (
	"fmt"
	"math/rand"
	"odev_four/lsm"
)

func main() {
	lsm := lsm.NewLSMTree(30, 30, "sstable_directory")

	fmt.Println("LSM Tree ve SSTable Kullanarak Veri Depolama Uygulaması")

	// Kullanıcı arayüzü ve işlevler
	for {
		fmt.Println("\n1. Ekle")
		fmt.Println("2. Getir")
		fmt.Println("3. Güncelle")
		fmt.Println("4. Sil")
		fmt.Println("5. Toplu Veri Ekle")
		fmt.Println("6. Toplu Veri Getir")
		fmt.Println("7. Çıkış")
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
			// Değer-Güncelle
			fmt.Print("Güncellemek istediğiniz anahtarı girin: ")
			var key int
			if _, err := fmt.Scan(&key); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
				continue
			}

			// Anahtarın var olup olmadığını kontrol et
			_, err := lsm.Get(key)
			if err != nil {
				fmt.Printf("Anahtar bulunamadı: %d\n", key)
				continue
			}

			fmt.Print("Yeni değeri girin: ")
			var newValue string
			if _, err := fmt.Scan(&newValue); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir metin girin.")
				continue
			}
			lsm.Put(key, newValue)
			fmt.Printf("Anahtar %d güncellendi.\n", key)
		case 4:
			fmt.Print("Silmek istediğiniz anahtarı girin: ")
			var key int
			if _, err := fmt.Scan(&key); err != nil {
				fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
				continue
			}
			lsm.Delete(key)
			fmt.Printf("Anahtar %d silindi.\n", key)
		case 5:
			BatchInsert(lsm)
			fmt.Printf("Toplu ekleme başarılı.\n")
		case 6:
			BatchGet(lsm)
			fmt.Printf("Toplu getirme başarılı.\n")
		case 7:
			fmt.Println("Programdan çıkılıyor.")
			return
		default:
			fmt.Println("Geçersiz seçenek. Lütfen 1 ile 4 arasında bir sayı girin.")
		}
	}
}

func BatchInsert(lsm *lsm.LSMTree) {
	// Veri ekleme işlemleri.
	for i := 1; i <= 60; i++ {
		key := i
		value := fmt.Sprintf("Value-%d", i)
		lsm.Put(key, value)
	}
}

func BatchGet(lsm *lsm.LSMTree) {
	// Veri okuma işlemi.
	for i := 1; i <= 10; i++ {
		key := rand.Intn(100) + 1
		value, err := lsm.Get(key)
		if err == nil {
			fmt.Printf("Key: %d, Value: %s\n", key, value)
		} else {
			fmt.Printf("Key not found: %d\n", key)
		}
	}
}
