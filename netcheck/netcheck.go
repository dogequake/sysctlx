package netcheck

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-ping/ping"
)

func Netcheck(host string) {
	fmt.Println("HOST:", host)   // Выводим введённый host
	fmt.Println(DNSLookup(host)) // Выводим результат DNS-резолва
	fmt.Println(PingHost(host))  // Выводим результат ping
	fmt.Println(HTTPCheck(host)) // Выводим результат HTTP-запроса
}

// DNSLookup выполняет DNS-резолв: возвращает первый найденный IPv4-адрес или сообщение об ошибке
func DNSLookup(host string) string {
	ips, err := net.LookupIP(host) // Получаем список IP-адресов по имени хоста
	if err != nil || len(ips) == 0 {
		return "IP: не найден"
	}
	for _, ip := range ips {
		if ip.To4() != nil { // Проверяем, что это IPv4
			return fmt.Sprintf("IP: %s", ip.String())
		}
	}
	return "IP: не найден"
}

// PingHost выполняет ICMP ping к хосту и возвращает результат (успех с временем или ошибку)
func PingHost(host string) string {
	pinger, err := ping.NewPinger(host) // Создаём новый пингер для хоста
	if err != nil {
		return "Ping: ошибка"
	}
	pinger.Count = 4   // Количество ICMP-запросов
	err = pinger.Run() // Запускаем пинг
	if err != nil {
		return "Ping: нет ответа"
	}
	stats := pinger.Statistics() // Получаем статистику
	return fmt.Sprintf("Ping: OK (%dмс)", stats.AvgRtt.Milliseconds())
}

// HTTPCheck выполняет HTTP GET-запрос к хосту (по протоколу http) и возвращает статус или ошибку
func HTTPCheck(host string) string {
	resp, err := http.Get("http://" + host) // Делаем HTTP-запрос
	if err != nil {
		return "HTTP: недоступен"
	}
	defer resp.Body.Close() // Не забываем закрыть тело ответа
	return fmt.Sprintf("HTTP: %s", resp.Status)
}
