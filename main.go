package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("proxies.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer inputFile.Close()

	file, err := os.Create("auth.lua")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(`local proxies = {
`)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "@")
		credentials := parts[0]
		address := parts[1]
		encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
		ipPort := strings.Split(address, ":")
		ip := ipPort[0]
		port := ipPort[1]
		writer.WriteString(fmt.Sprintf(`    { key = "%s", ip = "%s", port = "%s" },`, encoded, ip, port))
		writer.WriteString("\n")
	}

	writer.WriteString(`}

function set_proxy(txn)
    local index = (os.time() % #proxies) + 1
    local proxy = proxies[index]
    local auth = "Basic " .. proxy.key
    txn.http:req_set_header("Proxy-Authorization", auth)
    txn:set_var("txn.proxy_ip", proxy.ip)
    txn:set_var("txn.proxy_port", proxy.port)
end

core.register_action("set_proxy", { "http-req" }, set_proxy)
`)
	writer.Flush()
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}
}
