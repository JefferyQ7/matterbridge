// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SensitivityLabel undocumented
type SensitivityLabel struct {
	// Entity is the base model of SensitivityLabel
	Entity
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// ToolTip undocumented
	ToolTip *string `json:"toolTip,omitempty"`
	// IsEndpointProtectionEnabled undocumented
	IsEndpointProtectionEnabled *bool `json:"isEndpointProtectionEnabled,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// ApplicationMode undocumented
	ApplicationMode *ApplicationMode `json:"applicationMode,omitempty"`
	// LabelActions undocumented
	LabelActions []LabelActionBase `json:"labelActions,omitempty"`
	// AssignedPolicies undocumented
	AssignedPolicies []LabelPolicy `json:"assignedPolicies,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
	// AutoLabeling undocumented
	AutoLabeling *AutoLabeling `json:"autoLabeling,omitempty"`
	// Sublabels undocumented
	Sublabels []SensitivityLabel `json:"sublabels,omitempty"`
}