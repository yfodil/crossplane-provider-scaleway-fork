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

// ResolveReferences of this Device.
func (mg *Device) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HubIDRef,
		Selector:     mg.Spec.ForProvider.HubIDSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HubID")
	}
	mg.Spec.ForProvider.HubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HubIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.HubIDRef,
		Selector:     mg.Spec.InitProvider.HubIDSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HubID")
	}
	mg.Spec.InitProvider.HubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Network.
func (mg *Network) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HubIDRef,
		Selector:     mg.Spec.ForProvider.HubIDSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HubID")
	}
	mg.Spec.ForProvider.HubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HubIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.HubIDRef,
		Selector:     mg.Spec.InitProvider.HubIDSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HubID")
	}
	mg.Spec.InitProvider.HubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HubIDRef,
		Selector:     mg.Spec.ForProvider.HubIDSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HubID")
	}
	mg.Spec.ForProvider.HubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HubIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.HubIDRef,
		Selector:     mg.Spec.InitProvider.HubIDSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HubID")
	}
	mg.Spec.InitProvider.HubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HubIDRef = rsp.ResolvedReference

	return nil
}
