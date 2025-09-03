# sysctlx

Многофункциональный сетевой инструмент на Go: netcheck (ping, DNS, HTTP), portscan (сканер TCP-портов).

## Возможности

- Проверка доступности хоста: DNS, ICMP ping, HTTP
- Сканирование диапазона TCP-портов

## Установка

```sh
git clone https://github.com/dogequake/sysctlx.git
cd sysctlx
go build
```

## Использование

### Проверка хоста

```sh
./sysctlx netcheck google.com
```

### Сканирование портов

```sh
./sysctlx portscan google.com 20-25
```

## Пример вывода

```
HOST: google.com
IP: 142.250.74.238
Ping: OK (23мс)
HTTP: 200 OK
```

## Лицензия

MIT License