package google

import (
	"fmt"

	"github.com/cycloidio/inframap/errcode"
	"github.com/cycloidio/inframap/provider"

	tfdocGCP "github.com/cycloidio/tfdocs/providers/google"
	"github.com/cycloidio/tfdocs/resource"
)

// Provider is the Type that defines the interface 'provider.Interface'
type Provider struct {
	provider.NopProvider
}

// Type returns the type of the implementation
func (a Provider) Type() provider.Type { return provider.Google }

// IsNode is true if the resource is a Node
func (a Provider) IsNode(resource string) bool {
	_, ok := nodeTypes[resource]
	return ok
}

// IsEdge is true if the resource is an Edge
func (a Provider) IsEdge(resource string) bool {
	_, ok := edgeTypes[resource]
	return ok
}

// Resource returns the resource information
func (a Provider) Resource(resource string) (*resource.Resource, error) {
	r, err := tfdocGCP.GetResource(resource)
	if err != nil {
		return nil, fmt.Errorf("could not get resource %q: %w", resource, errcode.ErrProviderNotFoundResource)
	}
	return r, nil
}

// DataSource returns the resource information
func (a Provider) DataSource(resource string) (*resource.Resource, error) {
	r, err := tfdocGCP.GetDataSource(resource)
	if err != nil {
		return nil, fmt.Errorf("could not get resource %q: %w", resource, errcode.ErrProviderNotFoundDataSource)
	}
	return r, nil
}

// ResourceInOut returns the In and Out of the rs based on the cfg
func (a Provider) ResourceInOut(id, rs string, cfgs map[string]map[string]interface{}) ([]string, []string) {
	var tagins, tagouts []string
	cfg := cfgs[id]
	if rs == "google_compute_firewall" {
		if direction, ok := cfg["direction"]; ok {
			if direction.(string) == "INGRESS" {
				if tags, ok := cfg["target_tags"].([]interface{}); ok {
					for _, tag := range tags {
						tagouts = append(tagouts, tag.(string))
					}
				}
				if tags, ok := cfg["source_tags"].([]interface{}); ok {
					for _, tag := range tags {
						tagins = append(tagins, tag.(string))
					}
				}
			} else {
				if tags, ok := cfg["target_tags"].([]interface{}); ok {
					for _, tag := range tags {
						tagins = append(tagins, tag.(string))
					}
				}
			}
		}
	}

	// Once we know the tagins and tagouts we have to know
	// which resources have them to get the IDs of them
	var ins, outs []string
	for rsid, cfg := range cfgs {
		if id == rsid {
			// We ignore the resource we are in right now
			// because we want to know to which other resources
			// is it connected and not itself
			continue
		}
		pvid, ok := cfg["id"].(string)
		if !ok {
			// If for some reason it's not an String we continue
			continue
		}
		for _, in := range tagins {
			if tags, ok := cfg["tags"].([]interface{}); ok {
				if sliceInterfaceContains(tags, in) {
					ins = append(ins, pvid)
				}
			}
		}
		for _, out := range tagouts {
			if tags, ok := cfg["tags"].([]interface{}); ok {
				if sliceInterfaceContains(tags, out) {
					outs = append(outs, pvid)
				}
			}
		}
	}

	return ins, outs
}

func sliceInterfaceContains(is []interface{}, s interface{}) bool {
	for _, v := range is {
		if v == s {
			return true
		}
	}
	return false
}

// PreProcess will return extra edges to add them to the graph
// each element is an edge and for each edge we have the source
// and the target.
// [_][0] is the source of the edge
// [_][1] is the target of the edge
/*
step 1
======

we build a lookup table to store the firewalls ID associated to a tag

step 2
======

for each instance, we add to the edges result the relation between the tags and the list of
FW associated to this tags

*/
func (a Provider) PreProcess(cfg map[string]map[string]interface{}) [][]string {

	// fwTags will contain all the network tags present
	// and the firewall associated to it
	fwTags := make(map[string][]string, 0)

	edges := make([][]string, 0)

	// we build the lookup table
	for id, node := range cfg {
		if sources, ok := node["source_tags"]; ok {
			for _, source := range sources.([]interface{}) {
				fwTags[source.(string)] = append(fwTags[source.(string)], id)
			}
		}
		if targets, ok := node["target_tags"]; ok {
			for _, target := range targets.([]interface{}) {
				fwTags[target.(string)] = append(fwTags[target.(string)], id)
			}
		}
	}

	// add extra edges for each components with "tags" attributes
	for id, node := range cfg {
		if tags, ok := node["tags"]; ok {
			for _, tag := range tags.([]interface{}) {
				if fws, ok := fwTags[tag.(string)]; ok {
					for _, fw := range fws {
						edges = append(edges, []string{id, fw})
					}
				}
			}
		}
	}
	return edges
}
