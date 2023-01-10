package v1alpha2

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *KeptnAppVersion) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}
