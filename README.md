# LSM Tree ve SSTable Tabanlı Veri Depolama Uygulaması

Bu proje, LSM Tree (Log-Structured Merge-Tree) ve SSTable (Sorted Strings Table) kullanarak veri depolayan ve yöneten bir uygulamadır. Uygulama, key-value tabanlı veri saklama ve sorgulama işlevleri sunar. SSTable'lar "sstable_directory" dizini altında .txt formatında saklanır.

## Başlarken

Bu bölüm, projenin nasıl kurulacağı ve çalıştırılacağı hakkında adımları içerir.

### Önkoşullar

Projeyi çalıştırmak için sisteminizde Docker ve Docker Compose'un yüklü olması gerekmektedir. Docker'ın nasıl yükleneceği hakkında daha fazla bilgi için [Docker'ın resmi web sitesini](https://www.docker.com/get-started) ziyaret edebilirsiniz.

### Kurulum

Projeyi kurmak için aşağıdaki adımları izleyin:

1. Projeyi yerel makinenize klonlayın:
   ```bash
   git clone https://github.com/CosmosExmo/golang_simple_lsmtree.git
   ```
   Projeyi klonladıktan sonra, projenin root dizinine gidin:
   ```bash
   cd golang_simple_lsmtree
   ```

2. Projeyi Docker Compose ile çalıştırmak için aşağıdaki komutu kullanın:
   ```bash
   docker-compose run app
   ```
   Bu komut, uygulamanın Docker konteynerini interaktif modda başlatır ve kullanıcı arayüzüne erişmenizi sağlar.

## Kullanım

Uygulama başlatıldığında, terminalde bir menü görüntülenir. Bu menüden çeşitli işlemleri seçerek key-value tabanlı veri depolama ve sorgulama işlemlerini gerçekleştirebilirsiniz. Menüdeki seçenekler şunlardır:

- **Ekle:** Yeni bir anahtar-değer çifti ekler.
- **Getir:** Belirli bir anahtar için değeri getirir.
- **Güncelle:** Mevcut bir anahtar-değer çiftini günceller.
- **Sil:** Belirli bir anahtar için değeri siler.
- **Toplu Veri Ekle:** Birden fazla anahtar-değer çiftini toplu olarak ekler.
- **Toplu Veri Getir:** Birden fazla anahtarı sorgular.
- **Çıkış:** Uygulamadan çıkar.

## Lisans

Bu proje MIT Lisansı altında lisanslanmıştır.
