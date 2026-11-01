//ff:type feature=cli type=model
//ff:what dropError — a dump line read directly that does not decode as its entity: which entity and why
package main

// dropError says the dump line of entity ID does not decode (Err): the
// entity is left out (packResult.Dropped), as wikipg would have recorded a
// load_error and loaded no entity.
type dropError struct {
	ID  string
	Err error
}
