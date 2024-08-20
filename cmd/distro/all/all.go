package all

import (
	// The following are necessary as they register handlers in their init functions.

	_ "github.com/xmplusdev/xmcore/app/proxyman/inbound"
	_ "github.com/xmplusdev/xmcore/app/proxyman/outbound"

	// Required features. Can't remove unless there is replacements.
	// _ "github.com/xmplusdev/xmcore/app/dispatcher"
	_ "github.com/XMPlusDev/XMPlusv1/app/xdispatcher"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/xmplusdev/xmcore/app/commander"
	_ "github.com/xmplusdev/xmcore/app/log/command"
	_ "github.com/xmplusdev/xmcore/app/proxyman/command"
	_ "github.com/xmplusdev/xmcore/app/stats/command"

	// Other optional features.
	_ "github.com/xmplusdev/xmcore/app/dns"
	_ "github.com/xmplusdev/xmcore/app/log"
	_ "github.com/xmplusdev/xmcore/app/metrics"
	_ "github.com/xmplusdev/xmcore/app/policy"
	_ "github.com/xmplusdev/xmcore/app/reverse"
	_ "github.com/xmplusdev/xmcore/app/router"
	_ "github.com/xmplusdev/xmcore/app/stats"

	// Inbound and outbound proxies.
	_ "github.com/xmplusdev/xmcore/proxy/blackhole"
	_ "github.com/xmplusdev/xmcore/proxy/dns"
	_ "github.com/xmplusdev/xmcore/proxy/dokodemo"
	_ "github.com/xmplusdev/xmcore/proxy/freedom"
	_ "github.com/xmplusdev/xmcore/proxy/http"
	_ "github.com/xmplusdev/xmcore/proxy/shadowsocks"
	_ "github.com/xmplusdev/xmcore/proxy/socks"
	_ "github.com/xmplusdev/xmcore/proxy/trojan"
	_ "github.com/xmplusdev/xmcore/proxy/vless/inbound"
	_ "github.com/xmplusdev/xmcore/proxy/vless/outbound"
	_ "github.com/xmplusdev/xmcore/proxy/vmess/inbound"
	_ "github.com/xmplusdev/xmcore/proxy/vmess/outbound"
	_ "github.com/xmplusdev/xmcore/proxy/wireguard"

	// Transports
	_ "github.com/xmplusdev/xmcore/transport/internet/domainsocket"
	_ "github.com/xmplusdev/xmcore/transport/internet/http"
	_ "github.com/xmplusdev/xmcore/transport/internet/kcp"
	_ "github.com/xmplusdev/xmcore/transport/internet/quic"
	_ "github.com/xmplusdev/xmcore/transport/internet/tcp"
	_ "github.com/xmplusdev/xmcore/transport/internet/tls"
	_ "github.com/xmplusdev/xmcore/transport/internet/udp"
	_ "github.com/xmplusdev/xmcore/transport/internet/websocket"
	_ "github.com/xmplusdev/xmcore/transport/internet/reality"
	_ "github.com/xmplusdev/xmcore/transport/internet/httpupgrade"
	_ "github.com/xmplusdev/xmcore/transport/internet/splithttp"
	_ "github.com/xmplusdev/xmcore/transport/internet/grpc"

	// Transport headers
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/http"
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/noop"
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/srtp"
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/tls"
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/utp"
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/wechat"
	_ "github.com/xmplusdev/xmcore/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/xmplusdev/xmcore/main/json"
	_ "github.com/xmplusdev/xmcore/main/toml"
	_ "github.com/xmplusdev/xmcore/main/yaml"

	// Load config from file or http(s)
	_ "github.com/xmplusdev/xmcore/main/confloader/external"

	// Commands
	_ "github.com/xmplusdev/xmcore/main/commands/all"
)
