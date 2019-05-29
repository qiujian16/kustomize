// Code generated by pluginator on HashTransformer; DO NOT EDIT.
package builtin

import (
	"fmt"
	"sigs.k8s.io/kustomize/pkg/ifc"
	"sigs.k8s.io/kustomize/pkg/resmap"
)

type HashTransformerPlugin struct {
	hasher ifc.KunstructuredHasher
}

func NewHashTransformerPlugin() *HashTransformerPlugin {
	return &HashTransformerPlugin{}
}

func (p *HashTransformerPlugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, config []byte) (err error) {
	p.hasher = rf.RF().Hasher()
	return nil
}

// Transform appends hash to generated resources.
func (p *HashTransformerPlugin) Transform(m resmap.ResMap) error {
	for _, res := range m {
		if res.NeedHashSuffix() {
			h, err := p.hasher.Hash(res)
			if err != nil {
				return err
			}
			res.SetName(fmt.Sprintf("%s-%s", res.GetName(), h))
		}
	}
	return nil
}
