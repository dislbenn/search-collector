/*
IBM Confidential
OCO Source Materials
5737-E67
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/

package transforms

import (
	mcm "github.ibm.com/IBMPrivateCloud/hcm-compliance/pkg/apis/policy/v1alpha1"
)

type PolicyResource struct {
	*mcm.Policy
}

func (p PolicyResource) BuildNode() Node {
	node := transformCommon(p)         // Start off with the common properties
	apiGroupVersion(p.TypeMeta, &node) // add kind, apigroup and version
	// Extract the properties specific to this type
	node.Properties["remediationAction"] = string(p.Spec.RemediationAction)
	node.Properties["compliant"] = string(p.Status.ComplianceState)
	node.Properties["valid"] = p.Status.Valid

	rules := int64(0)
	if p.Spec.RoleTemplates != nil {
		for _, role := range p.Spec.RoleTemplates {
			if role != nil {
				rules += int64(len(role.Rules))
			}
		}
	}
	node.Properties["numRules"] = rules
	pnamespace, ok := p.ObjectMeta.Labels["parent-namespace"]
	if ok {
		node.Properties["parent-namespace"] = pnamespace
	}
	ppolicy, ok := p.ObjectMeta.Labels["parent-policy"]
	if ok {
		node.Properties["parent-policy"] = ppolicy
	}

	return node
}

func (p PolicyResource) BuildEdges(ns NodeStore) []Edge {
	//no op for now to implement interface
	return []Edge{}
}
