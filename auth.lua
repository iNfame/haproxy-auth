local proxies = {
    { key = "a0FZQXYxOmh5UlV2OHNBTmFQWg==", ip = "cproxy.site", port = "12338" },
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
