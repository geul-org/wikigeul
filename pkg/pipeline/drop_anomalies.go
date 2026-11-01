//ff:func feature=pipeline type=command control=sequence
//ff:what dropAnomalies — a dump line that does not decode as an anomaly, unless -allow accepts it
package pipeline

// dropReason is the reason of a decode anomaly, which -allow names as
// "decode:덤프 줄 해석 실패".
const dropReason = "덤프 줄 해석 실패"

// dropAnomalies is the anomaly of entity id whose dump line does not decode
// (why, the error), class 데이터 이상; none when allow has "decode:" +
// dropReason.
func dropAnomalies(id, why string, allow map[string]bool) []anomaly {
	if allow["decode:"+dropReason] {
		return nil
	}
	return []anomaly{{ID: id, Kind: "decode", Value: why, Reason: dropReason, Class: "데이터 이상"}}
}
