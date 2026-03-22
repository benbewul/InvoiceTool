# Fatura Mutabakat Ekranı

PCC ve BPPS sistemlerinde fatura sorgulama ve karşılaştırma yapan demo arayüz.

## Demo fatura numaraları
- FD60351N9DD0EA
- FD60351NA6DA85
- FD60351O9AB618

## Çalıştırma
```bash
go run ./cmd/server
```

Tarayıcı:
`http://localhost:8080`

## Docker
```bash
docker build -t fatura-mutabakat-ekrani .
docker run -p 8080:8080 fatura-mutabakat-ekrani
```
