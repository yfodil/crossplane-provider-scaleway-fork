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

// ResolveReferences of this Container.
func (mg *Container) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ContainerNamespaceList{},
			Managed: &ContainerNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NamespaceIDRef,
		Selector:     mg.Spec.InitProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ContainerNamespaceList{},
			Managed: &ContainerNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NamespaceID")
	}
	mg.Spec.InitProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Cron.
func (mg *Cron) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContainerIDRef,
		Selector:     mg.Spec.ForProvider.ContainerIDSelector,
		To: reference.To{
			List:    &ContainerList{},
			Managed: &Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerID")
	}
	mg.Spec.ForProvider.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContainerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ContainerIDRef,
		Selector:     mg.Spec.InitProvider.ContainerIDSelector,
		To: reference.To{
			List:    &ContainerList{},
			Managed: &Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ContainerID")
	}
	mg.Spec.InitProvider.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ContainerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContainerIDRef,
		Selector:     mg.Spec.ForProvider.ContainerIDSelector,
		To: reference.To{
			List:    &ContainerList{},
			Managed: &Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerID")
	}
	mg.Spec.ForProvider.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContainerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ContainerIDRef,
		Selector:     mg.Spec.InitProvider.ContainerIDSelector,
		To: reference.To{
			List:    &ContainerList{},
			Managed: &Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ContainerID")
	}
	mg.Spec.InitProvider.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ContainerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Token.
func (mg *Token) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContainerIDRef,
		Selector:     mg.Spec.ForProvider.ContainerIDSelector,
		To: reference.To{
			List:    &ContainerList{},
			Managed: &Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerID")
	}
	mg.Spec.ForProvider.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ContainerNamespaceList{},
			Managed: &ContainerNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContainerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ContainerIDRef,
		Selector:     mg.Spec.InitProvider.ContainerIDSelector,
		To: reference.To{
			List:    &ContainerList{},
			Managed: &Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ContainerID")
	}
	mg.Spec.InitProvider.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ContainerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NamespaceIDRef,
		Selector:     mg.Spec.InitProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ContainerNamespaceList{},
			Managed: &ContainerNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NamespaceID")
	}
	mg.Spec.InitProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}
