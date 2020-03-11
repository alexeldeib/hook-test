/*
Copyright 2020 Ace Eldeib.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	shipv1beta1 "github.com/alexeldeib/test-hook/api/v1beta1"
)

// FrigateReconciler reconciles a Frigate object
type FrigateReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=ship.example.org,resources=frigates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ship.example.org,resources=frigates/status,verbs=get;update;patch

func (r *FrigateReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("frigate", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *FrigateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&shipv1beta1.Frigate{}).
		Complete(r)
}
