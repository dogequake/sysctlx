// Программа для проверки доступности хоста:
// 1. Выполняет DNS-резолв (получает IP-адрес по имени хоста)
// 2. Пингует хост (ICMP ping)
// 3. Проверяет HTTP-доступность (делает HTTP-запрос)
// Выводит результат в виде:
// HOST: <host>
// IP: <ip или сообщение об ошибке>
// Ping: OK (<время>) или сообщение об ошибке
// HTTP: <код ответа или сообщение об ошибке>
//
// Использование: go run main.go, затем ввести host с клавиатуры
//
// Для работы ICMP ping требуется root-права или разрешение на raw-сокеты.

package main

// main — точка входа. Запрашивает у пользователя host и выводит результаты всех проверок.
import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sysctlx/netcheck"
	"sysctlx/portscan"
)

func main() {
	args := os.Args
	if len(args) == 3 && args[1] == "netcheck" {
		netcheck.Netcheck(args[2])
		return
	}
	if len(args) == 4 && args[1] == "portscan" {
		host := args[2]
		ports, err := parsePortRange(args[3])
		if err != nil {
			fmt.Println("Ошибка диапазона портов. Пример: 20-1024")
			return
		}
		portscan.Portscan(host, ports)
		return
	}
	fmt.Println("Использование:")
	fmt.Println("  sysctlx netcheck <host>")
	fmt.Println("  sysctlx portscan <host> <start-end>")
}

// parsePortRange("20-25") -> []int{20,21,22,23,24,25}
func parsePortRange(s string) ([]int, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("bad range")
	}
	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || start > end {
		return nil, fmt.Errorf("bad range")
	}
	ports := make([]int, 0, end-start+1)
	for p := start; p <= end; p++ {
		ports = append(ports, p)
	}
	return ports, nil
}
