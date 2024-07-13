# XMPlus
XMPlus v2.0 backend


#### Config directory
```
cd /etc/XMPlus
```

### Onclick XMPlus backennd Install
```
bash <(curl -Ls https://raw.githubusercontent.com/XMPlusDev/XMPlusv1/install/install.sh)
```

### /etc/XMPlus/config.yml
```
Log:
  Level: warning # Log level: none, error, warning, info, debug 
  AccessPath: # /etc/XMPlus/access.Log
  ErrorPath: # /etc/XMPlus/error.log
DnsConfigPath:  #/etc/XMPlus/dns.json
RouteConfigPath: # /etc/XMPlus/route.json
InboundConfigPath: # /etc/XMPlus/inbound.json
OutboundConfigPath: # /etc/XMPlus/outbound.json
ConnectionConfig:
  Handshake: 8 
  ConnIdle: 300 
  UplinkOnly: 0 
  DownlinkOnly: 0 
  BufferSize: 64
Nodes:
  -
    ApiConfig:
      ApiHost: "https://www.xyz.com"
      ApiKey: "123"
      NodeID: 1
      Timeout: 30 
      RuleListPath: # /etc/XMPlus/rulelist Path to local rulelist file
    ControllerConfig:
      EnableDNS: false # Use custom DNS config, Please ensure that you set the dns.json well
      DNSStrategy: AsIs # AsIs, UseIP, UseIPv4, UseIPv6
      CertConfig:
        Email: author@xmplus.dev                    # Required when Cert Mode is not none
        CertFile: /etc/XMPlus/node1.xmplus.dev.crt  # Required when Cert Mode is file
        KeyFile: /etc/XMPlus/node1.xmplus.dev.key   # Required when Cert Mode is file
        Provider: cloudflare                        # Required when Cert Mode is dns
        CertEnv:                                    # Required when Cert Mode is dns
          CLOUDFLARE_EMAIL:                         # Required when Cert Mode is dns
          CLOUDFLARE_API_KEY:                       # Required when Cert Mode is dns
      EnableFallback: false # Only support for Trojan and Vless
      FallBackConfigs:  # Support multiple fallbacks
        - SNI: # TLS SNI(Server Name Indication), Empty for any
          Alpn: # Alpn, Empty for any
          Path: # HTTP PATH, Empty for any
          Dest: 80 # Required, Destination of fallback, check https://xtls.github.io/config/features/fallback.html for details.
          ProxyProtocolVer: 0 # Send PROXY protocol version, 0 for disable
```

## XMPlus Panel Server configuration

### Network Settings

#### TCP
```
{
  "transport" : "tcp",
  "acceptProxyProtocol": false,
  "flow": "xtls-rprx-vision",
  "header": {
    "type": "none"
  }
}
```
#### TCP + HTTP
```
{
  "transport" : "tcp",
  "acceptProxyProtocol": false,
  "header": {
    "type": "http",
    "request": {
      "path": "/xmplus",
      "headers": {
        "Host": ["www.baidu.com", "www.taobao.com", "www.cloudflare.com"]
      }
    }
  }
}
```
####  WS
```
{
  "transport": "ws",
  "acceptProxyProtocol": false,
  "path": "/xmplus?ed=2560",
  "host": "hk1.xyz.com",
  "headers": {
    "host": "hk1.xyz.com"
  }
}
```

####  H2
```
{
  "transport" : "h2",
  "acceptProxyProtocol": false,
  "host": "hk1.xyz.com",
  "path": "/"
}
```

####  GRPC
```
{
  "transport" : "grpc",
  "acceptProxyProtocol": false,
  "serviceName": "xmplus",
  "authority": "hk1.xyz.com"
}
```

####  HTTPUPGRADE
```
{
  "transport" : "httpupgrade",
  "acceptProxyProtocol": false,
  "host": "hk1.xyz.com",
  "path": "/xmplus?ed=2560",
  "headers": {
    "host": "hk1.xyz.com"
  }
}
```

####  SPLITHTTP
```
{
  "transport" : "splithttp",
  "host": "hk1.xyz.com",
  "path": "/",
  "headers": {
    "host": "hk1.xyz.com"
  },
  "maxUploadSize": 1000000,
  "maxConcurrentUploads": 10 
}
```

####  QUIC
```
{
  "transport" : "quic",
  "acceptProxyProtocol": false,
  "security": "none",
  "key": "",
  "header": {
    "type": "none"
  }
}
```
####  KCP
```
{
  "transport" : "kcp",
  "acceptProxyProtocol": false,
  "congestion": false,
  "header": {
    "type": "none"
  },
  "seed": "password"
}
```

### Security Settings


#### TLS
```
{
  "serverName": "xmplus.dev",
  "rejectUnknownSni": true,
  "allowInsecure": false,
  "fingerprint": "chrome",
  "sni" : "xmplus.dev"
}
```
#### REALITY

[Generate Private and Public Keys Here](https://go.dev/play/p/N5kQhIjtye7)

```
{
  "show" : false,
  "dest": "www.lovelive-anime.jp:443",
  "privatekey" : "yBaw532IIUNuQWDTncozoBaLJmcd1JZzvsHUgVPxMk8",
  "minclientver":"",
  "maxclientver":"",
  "maxtimediff":0,
  "proxyprotocol":0,
  "shortids" : [
    "6ba85179e30d4fc2"
  ],
  "serverNames": [
    "www.lovelive-anime.jp",
    "www.cloudflare.com"
  ],
  "fingerprint": "chrome",
  "spiderx": "",
  "publickey": "7xhH4b_VkliBxGulljcyPOH-bYUA2dl-XAdZAsfhk04"
}
```
