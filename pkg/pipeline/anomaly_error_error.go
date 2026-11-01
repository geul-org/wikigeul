//ff:func feature=pipeline type=helper control=sequence
//ff:what AnomalyError.Error — the stop message: entity, count, report path
package pipeline

import "fmt"

// Error says where the strict run stopped and where the report is.
func (e AnomalyError) Error() string {
	return fmt.Sprintf("wikigeul -strict: %d anomalies at %s, stopped after a checkpoint; report %s", e.Count, e.ID, e.Path)
}
