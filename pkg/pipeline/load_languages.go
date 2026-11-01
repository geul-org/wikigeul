//ff:func feature=pipeline type=loader control=sequence
//ff:what LoadLanguages — read the GEUL Language Registry once for the run into the config
package pipeline

// LoadLanguages opens the archive at c.LangRoot (registryDir) and reads
// version c.LangVer of the registry into c.Langs and c.LangDoc
// (languageRegistry).
func LoadLanguages(c *Config) error {
	reg, err := registryDir(c.LangRoot)
	if err != nil {
		return err
	}
	c.Langs, c.LangDoc, err = languageRegistry(reg, c.LangVer)
	return err
}
