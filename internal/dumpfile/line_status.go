//ff:type feature=dumpfile type=model
//ff:what lineStatus — what a dump line is to the sort: no entity, kept, left out by the filter, or not readable
package dumpfile

// lineStatus is what lineRecord found on a dump line.
type lineStatus int

const (
	lineNone    lineStatus = iota // "[", "]" or blank: no entity
	lineKept                      // an entity the filter keeps
	lineSkipped                   // an entity the filter leaves out (another kind, range, list)
	lineBad                       // no readable id, or an ID wikipg cannot split
)
