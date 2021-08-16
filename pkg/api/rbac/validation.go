/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// from ./validation/validation.go
package rbac

import (
	path "github.com/yubo/apiserver/pkg/api/validation"
	"github.com/yubo/golib/util/validation/field"
)

// ValidateRBACName is exported to allow types outside of the RBAC API group to reuse this validation logic
// Minimal validation of names for roles and bindings. Identical to the validation for Openshift. See:
// * https://github.com/kubernetes/kubernetes/blob/60db50/pkg/api/validation/name.go
// * https://github.com/openshift/origin/blob/388478/pkg/api/helpers.go
func ValidateRBACName(name string, prefix bool) []string {
	return path.IsValidPathSegmentName(name)
}

func ValidateRole(role *Role) field.ErrorList {
	allErrs := field.ErrorList{}

	for i, rule := range role.Rules {
		if err := ValidatePolicyRule(rule, true, field.NewPath("rules").Index(i)); err != nil {
			allErrs = append(allErrs, err...)
		}
	}
	if len(allErrs) != 0 {
		return allErrs
	}
	return nil
}

func ValidateRoleUpdate(role *Role, oldRole *Role) field.ErrorList {
	allErrs := ValidateRole(role)

	return allErrs
}

func ValidateClusterRole(role *ClusterRole) field.ErrorList {
	return nil
}

func ValidateClusterRoleUpdate(role *ClusterRole, oldRole *ClusterRole) field.ErrorList {
	allErrs := ValidateClusterRole(role)

	return allErrs
}

// ValidatePolicyRule is exported to allow types outside of the RBAC API group to embed a PolicyRule and reuse this validation logic
func ValidatePolicyRule(rule PolicyRule, isNamespaced bool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if len(rule.Verbs) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("verbs"), "verbs must contain at least one value"))
	}

	if len(rule.NonResourceURLs) > 0 {
		if isNamespaced {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("nonResourceURLs"), rule.NonResourceURLs, "namespaced rules cannot apply to non-resource URLs"))
		}
		if len(rule.Resources) > 0 || len(rule.ResourceNames) > 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("nonResourceURLs"), rule.NonResourceURLs, "rules cannot apply to both regular resources and non-resource URLs"))
		}
		return allErrs
	}

	//if len(rule.APIGroups) == 0 {
	//	allErrs = append(allErrs, field.Required(fldPath.Child("apiGroups"), "resource rules must supply at least one api group"))
	//}
	if len(rule.Resources) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("resources"), "resource rules must supply at least one resource"))
	}
	return allErrs
}

func ValidateRoleBinding(roleBinding *RoleBinding) field.ErrorList {
	allErrs := field.ErrorList{}

	// TODO allow multiple API groups.  For now, restrict to one, but I can envision other experimental roles in other groups taking
	// advantage of the binding infrastructure
	//if roleBinding.RoleRef.APIGroup != GroupName {
	//	allErrs = append(allErrs, field.NotSupported(field.NewPath("roleRef", "apiGroup"), roleBinding.RoleRef.APIGroup, []string{GroupName}))
	//}

	switch roleBinding.RoleRef.Kind {
	case "Role", "ClusterRole":
	default:
		allErrs = append(allErrs, field.NotSupported(field.NewPath("roleRef", "kind"), roleBinding.RoleRef.Kind, []string{"Role", "ClusterRole"}))

	}

	if len(roleBinding.RoleRef.Name) == 0 {
		allErrs = append(allErrs, field.Required(field.NewPath("roleRef", "name"), ""))
	} else {
		for _, msg := range ValidateRBACName(roleBinding.RoleRef.Name, false) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("roleRef", "name"), roleBinding.RoleRef.Name, msg))
		}
	}

	subjectsPath := field.NewPath("subjects")
	for i, subject := range roleBinding.Subjects {
		allErrs = append(allErrs, ValidateRoleBindingSubject(subject, true, subjectsPath.Index(i))...)
	}

	return allErrs
}

func ValidateRoleBindingUpdate(roleBinding *RoleBinding, oldRoleBinding *RoleBinding) field.ErrorList {
	allErrs := ValidateRoleBinding(roleBinding)

	if oldRoleBinding.RoleRef != roleBinding.RoleRef {
		allErrs = append(allErrs, field.Invalid(field.NewPath("roleRef"), roleBinding.RoleRef, "cannot change roleRef"))
	}

	return allErrs
}

func ValidateClusterRoleBinding(roleBinding *ClusterRoleBinding) field.ErrorList {
	allErrs := field.ErrorList{}

	// TODO allow multiple API groups.  For now, restrict to one, but I can envision other experimental roles in other groups taking
	// advantage of the binding infrastructure
	//if roleBinding.RoleRef.APIGroup != GroupName {
	//	allErrs = append(allErrs, field.NotSupported(field.NewPath("roleRef", "apiGroup"), roleBinding.RoleRef.APIGroup, []string{GroupName}))
	//}

	switch roleBinding.RoleRef.Kind {
	case "ClusterRole":
	default:
		allErrs = append(allErrs, field.NotSupported(field.NewPath("roleRef", "kind"), roleBinding.RoleRef.Kind, []string{"ClusterRole"}))

	}

	if len(roleBinding.RoleRef.Name) == 0 {
		allErrs = append(allErrs, field.Required(field.NewPath("roleRef", "name"), ""))
	} else {
		for _, msg := range ValidateRBACName(roleBinding.RoleRef.Name, false) {
			allErrs = append(allErrs, field.Invalid(field.NewPath("roleRef", "name"), roleBinding.RoleRef.Name, msg))
		}
	}

	subjectsPath := field.NewPath("subjects")
	for i, subject := range roleBinding.Subjects {
		allErrs = append(allErrs, ValidateRoleBindingSubject(subject, false, subjectsPath.Index(i))...)
	}

	return allErrs
}

func ValidateClusterRoleBindingUpdate(roleBinding *ClusterRoleBinding, oldRoleBinding *ClusterRoleBinding) field.ErrorList {
	allErrs := ValidateClusterRoleBinding(roleBinding)

	if oldRoleBinding.RoleRef != roleBinding.RoleRef {
		allErrs = append(allErrs, field.Invalid(field.NewPath("roleRef"), roleBinding.RoleRef, "cannot change roleRef"))
	}

	return allErrs
}

// ValidateRoleBindingSubject is exported to allow types outside of the RBAC API group to embed a Subject and reuse this validation logic
func ValidateRoleBindingSubject(subject Subject, isNamespaced bool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(subject.Name) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
	}

	switch subject.Kind {
	case ServiceAccountKind:
		//if len(subject.APIGroup) > 0 {
		//	allErrs = append(allErrs, field.NotSupported(fldPath.Child("apiGroup"), subject.APIGroup, []string{""}))
		//}
		if !isNamespaced && len(subject.Namespace) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("namespace"), ""))
		}

	case UserKind:
		// TODO(ericchiang): What other restrictions on user name are there?
		//if subject.APIGroup != GroupName {
		//	allErrs = append(allErrs, field.NotSupported(fldPath.Child("apiGroup"), subject.APIGroup, []string{GroupName}))
		//}

	case GroupKind:
		// TODO(ericchiang): What other restrictions on group name are there?
		//if subject.APIGroup != GroupName {
		//	allErrs = append(allErrs, field.NotSupported(fldPath.Child("apiGroup"), subject.APIGroup, []string{GroupName}))
		//}

	default:
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("kind"), subject.Kind, []string{ServiceAccountKind, UserKind, GroupKind}))
	}

	return allErrs
}
