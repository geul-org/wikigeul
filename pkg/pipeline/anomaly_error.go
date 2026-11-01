//ff:type feature=pipeline type=model
//ff:what AnomalyError — a strict run stopped at anomalies: the first entity, how many, where the report is
package pipeline

// AnomalyError is returned when a strict run stops (packWriter.halt); main
// exits 3 on it.
type AnomalyError struct {
	ID    string
	Count int
	Path  string
}
