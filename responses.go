package smitego

import "fmt"

type createSessionResp struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}

type DataUsed struct {
	ActiveSessions     int64  `json:"Active_Session"`
	ConcurrentSessions int64  `json:"Concurrent_Sessions"`
	RequestLimitDaily  int64  `json:"Request_Limit_Daily"`
	SessionCap         int64  `json:"Session_Cap"`
	SessionTimeLimit   int64  `json:"Session_Time_Limit"`
	TotalRequestsToday int64  `json:"Total_Requests_Today"`
	TotalSessionsToday int64  `json:"Total_Sessions_Today"`
	RetMsg             string `json:"ret_msg"`
}

func (d *DataUsed) String() string {
	return fmt.Sprintf("Active: %d Concurrent: %d Limit: %d Session Cap: %d Time Limit: %d Today's requests: %d Today's session %d RetMsg %s", d.ActiveSessions, d.ConcurrentSessions, d.RequestLimitDaily, d.SessionCap, d.SessionTimeLimit, d.TotalRequestsToday, d.TotalSessionsToday, d.RetMsg)
}
