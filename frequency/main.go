package main

import "github.com/rwirdemann/legacylab/git"

func main() {
	git.Checkout("https://devstack.vwgroup.com/bitbucket/scm/ngw/vwg.idhub.core-ds.git")
	git.ChangeFrequency("/Users/ralf/tmp/vwg.idhub.core-ds", 30)
}
