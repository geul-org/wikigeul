//ff:type feature=pipeline type=model
//ff:what DropError — a dump line read directly that does not decode as its entity: which entity and why
package pipeline

// DropError says the dump line of entity ID does not decode (Err): the
// entity is left out (packResult.Dropped), as wikipgeul would have recorded a
// load_error and loaded no entity.
type DropError struct {
	ID  string
	Err error
}
