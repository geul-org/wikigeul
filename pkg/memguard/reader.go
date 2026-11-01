//ff:type feature=memguard type=model
//ff:what Reader — where samples come from; given the phase the run is in, so a test can press on one phase
package memguard

// Reader takes one Sample. The system's reader (procReader) ignores phase.
type Reader func(phase string) (Sample, error)
