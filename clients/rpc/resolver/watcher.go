package resolver

import "time"

// 地址更新
func (m mcenterResolver) Watch() {
	for {
		time.Sleep(m.refreshTime())
		m.resolve()
	}
}
