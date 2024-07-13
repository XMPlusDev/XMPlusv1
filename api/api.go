// To implement an api , one needs to implement the interface below.

package api

type API interface {
	GetNodeInfo() (nodeInfo *NodeInfo, err error)
	GetRelayNodeInfo() (nodeInfo *RelayNodeInfo, err error)
	GetUserList() (userList *[]UserInfo, err error)
	ReportNodeOnlineIPs(onlineIP *[]OnlineIP) (err error)
	ReportUserTraffic(userTraffic *[]UserTraffic) (err error)
	Describe() ClientInfo
	GetNodeRule() (ruleList *[]DetectRule, err error)
	Debug()
}
