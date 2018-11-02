package v1
import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionMeshFunc func(original, desired *Mesh) (bool, error)

type MeshReconciler interface {
	Reconcile(namespace string, desiredResources MeshList, transition TransitionMeshFunc, opts clients.ListOpts) error
}

func meshsToResources(list MeshList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, mesh := range list {
		resourceList = append(resourceList, mesh)
	}
	return resourceList
}

func NewMeshReconciler(client MeshClient) MeshReconciler {
	return &meshReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type meshReconciler struct {
	base reconcile.Reconciler
}

func (r *meshReconciler) Reconcile(namespace string, desiredResources MeshList, transition TransitionMeshFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "mesh_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Mesh), desired.(*Mesh))
		}
	}
	return r.base.Reconcile(namespace, meshsToResources(desiredResources), transitionResources, opts)
}
