local proxies = {
    { key = "dXNlcjkyNjIyOnd4MHQ2bQ==", ip = "149.126.246.46", port = "5126" },
    { key = "dXNlcjkyNjIyOnd4MHQ2bQ==", ip = "149.126.246.199", port = "5126" },
    { key = "QVVEYk5kMUs6MTVheWkyVmo=", ip = "45.140.62.181", port = "62192" },
}

function set_proxy(txn)
    local index = (os.time() % #proxies) + 1
    local proxy = proxies[index]
    local auth = "Basic " .. proxy.key
    txn.http:req_set_header("Proxy-Authorization", auth)
    txn:set_var("txn.proxy_ip", proxy.ip)
    txn:set_var("txn.proxy_port", proxy.port)
end

core.register_action("set_proxy", { "http-req" }, set_proxy)
