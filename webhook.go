package main

import (
	"context"

	apiv1beta1 "github.com/alexeldeib/test-hook/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var _ inject.Client = &MyTypeValidator{}

type MyTypeValidator struct {
	client.Client
	*admission.Decoder
}

func (v *MyTypeValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	_ = &apiv1beta1.Frigate{}

	return admission.Allowed("")
}

// MyTypeValidator implements inject.Client.
// A client will be automatically injected.

// InjectClient injects the client.
func (v *MyTypeValidator) InjectClient(c client.Client) error {
	v.Client = c
	return nil
}

// MyTypeValidator implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (v *MyTypeValidator) InjectDecoder(d *admission.Decoder) error {
	v.Decoder = d
	return nil
}
