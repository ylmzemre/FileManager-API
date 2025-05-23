# FileManager-API
## Kısa açıklama
  Backend : Go + Gin + GORM + JWT  
  Frontend : Basit HTML/JS tek sayfa (static/index.html)  
  PDF / PNG / JPG dosyalarını yükleme-listeleme-silme.  
  Dosya gövdesi disk’te /uploads, meta‐veri PostgreSQL’de tutulur.  

## Gerekenler
  Go 1.22+  
  PostgreSQL 14+ (yerel veya Docker)  
  Git (klonlamak için)  

## Kurulum
Kaynak kodu alın  
```
git clone https://github.com/<kullanıcı>/file-manager.git
cd file-manager
```
Database Kurulumu (Docker ile)  
```
docker run --name file-db \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=file_db \
  -p 5432:5432 -d postgres:15-alpine
```
Ortam Değişkenlerini Ayarlayın  
```
export DB_DSN="host=localhost user=postgres password=postgres dbname=file_db port=5432 sslmode=disable"
export JWT_SECRET="supersecret"      # rastgele bir ifade
export UPLOAD_DIR="./uploads"
```
Bağımlılıkları Kur, Çalıştır
```
go mod tidy
go run main.go
```
