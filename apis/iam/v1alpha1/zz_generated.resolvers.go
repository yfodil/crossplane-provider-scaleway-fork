/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ApiKey.
func (mg *ApiKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ApplicationIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.ApplicationIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ApplicationIdsRefs,
		Selector:      mg.Spec.ForProvider.ApplicationIdsSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationIds")
	}
	mg.Spec.ForProvider.ApplicationIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ApplicationIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.ApplicationIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.ApplicationIdsRefs,
		Selector:      mg.Spec.InitProvider.ApplicationIdsSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationIds")
	}
	mg.Spec.InitProvider.ApplicationIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ApplicationIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ApplicationIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.GroupIDRef,
		Selector:     mg.Spec.InitProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupID")
	}
	mg.Spec.InitProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupIDRef = rsp.ResolvedReference

	return nil
}
