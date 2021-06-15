// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1




import (
    "context"

    "github.com/solo-io/skv2/pkg/controllerutils"
    "github.com/solo-io/skv2/pkg/multicluster"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/client-go/kubernetes/scheme"
    "k8s.io/client-go/rest"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the settings.mesh.gloo.solo.io/v1 APIs
type MulticlusterClientset interface {
    // Cluster returns a Clientset for the given cluster
    Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
    client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
    return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
    client, err := m.client.Cluster(cluster)
    if err != nil {
        return nil, err
    }
    return NewClientset(client), nil
}

// clienset for the settings.mesh.gloo.solo.io/v1 APIs
type Clientset interface {
    // clienset for the settings.mesh.gloo.solo.io/v1/v1 APIs
    Settings() SettingsClient
    // clienset for the settings.mesh.gloo.solo.io/v1/v1 APIs
    Dashboards() DashboardClient
}

type clientSet struct {
    client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
    scheme := scheme.Scheme
    if err := AddToScheme(scheme); err != nil{
        return nil, err
    }
    client, err := client.New(cfg, client.Options{
        Scheme: scheme,
    })
    if err != nil {
        return nil, err
    }
    return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
    return &clientSet{client: client}
}

// clienset for the settings.mesh.gloo.solo.io/v1/v1 APIs
func (c *clientSet) Settings() SettingsClient {
    return NewSettingsClient(c.client)
}

// clienset for the settings.mesh.gloo.solo.io/v1/v1 APIs
func (c *clientSet) Dashboards() DashboardClient {
    return NewDashboardClient(c.client)
}

// Reader knows how to read and list Settingss.
type SettingsReader interface {
    // Get retrieves a Settings for the given object key
    GetSettings(ctx context.Context, key client.ObjectKey) (*Settings, error)

    // List retrieves list of Settingss for a given namespace and list options.
    ListSettings(ctx context.Context, opts ...client.ListOption) (*SettingsList, error)
}

// SettingsTransitionFunction instructs the SettingsWriter how to transition between an existing
// Settings object and a desired on an Upsert
type SettingsTransitionFunction func(existing, desired *Settings) error

// Writer knows how to create, delete, and update Settingss.
type SettingsWriter interface {
    // Create saves the Settings object.
    CreateSettings(ctx context.Context, obj *Settings, opts ...client.CreateOption) error

    // Delete deletes the Settings object.
    DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

    // Update updates the given Settings object.
    UpdateSettings(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

    // Patch patches the given Settings object.
    PatchSettings(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error

    // DeleteAllOf deletes all Settings objects matching the given options.
    DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error

    // Create or Update the Settings object.
    UpsertSettings(ctx context.Context, obj *Settings, transitionFuncs ...SettingsTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Settings object.
type SettingsStatusWriter interface {
    // Update updates the fields corresponding to the status subresource for the
    // given Settings object.
    UpdateSettingsStatus(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

    // Patch patches the given Settings object's subresource.
    PatchSettingsStatus(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Settingss.
type SettingsClient interface {
    SettingsReader
    SettingsWriter
    SettingsStatusWriter
}

type settingsClient struct {
    client client.Client
}

func NewSettingsClient(client client.Client) *settingsClient {
    return &settingsClient{client: client}
}


func (c *settingsClient) GetSettings(ctx context.Context, key client.ObjectKey) (*Settings, error) {
    obj := &Settings{}
    if err := c.client.Get(ctx, key, obj); err != nil {
        return nil, err
    }
    return obj, nil
}

func (c *settingsClient) ListSettings(ctx context.Context, opts ...client.ListOption) (*SettingsList, error) {
    list := &SettingsList{}
    if err := c.client.List(ctx, list, opts...); err != nil {
        return nil, err
    }
    return list, nil
}

func (c *settingsClient) CreateSettings(ctx context.Context, obj *Settings, opts ...client.CreateOption) error {
    return c.client.Create(ctx, obj, opts...)
}


func (c *settingsClient) DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
    obj := &Settings{}
    obj.SetName(key.Name)
    obj.SetNamespace(key.Namespace)
    return c.client.Delete(ctx, obj, opts...)
}

func (c *settingsClient) UpdateSettings(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
    return c.client.Update(ctx, obj, opts...)
}

func (c *settingsClient) PatchSettings(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error {
    return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *settingsClient) DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error {
    obj := &Settings{}
    return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *settingsClient) UpsertSettings(ctx context.Context, obj *Settings, transitionFuncs ...SettingsTransitionFunction) error {
    genericTxFunc := func(existing, desired runtime.Object) error {
        for _, txFunc := range transitionFuncs {
            if err := txFunc(existing.(*Settings), desired.(*Settings)); err != nil {
                return err
            }
        }
        return nil
    }
    _, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
    return err
}

func (c *settingsClient) UpdateSettingsStatus(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
    return c.client.Status().Update(ctx, obj, opts...)
}

func (c *settingsClient) PatchSettingsStatus(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error {
    return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides SettingsClients for multiple clusters.
type MulticlusterSettingsClient interface {
    // Cluster returns a SettingsClient for the given cluster
    Cluster(cluster string) (SettingsClient, error)
}

type multiclusterSettingsClient struct {
    client multicluster.Client
}

func NewMulticlusterSettingsClient(client multicluster.Client) MulticlusterSettingsClient {
    return &multiclusterSettingsClient{client: client}
}

func (m *multiclusterSettingsClient) Cluster(cluster string) (SettingsClient, error) {
    client, err := m.client.Cluster(cluster)
    if err != nil {
        return nil, err
    }
    return  NewSettingsClient(client), nil
}

// Reader knows how to read and list Dashboards.
type DashboardReader interface {
    // Get retrieves a Dashboard for the given object key
    GetDashboard(ctx context.Context, key client.ObjectKey) (*Dashboard, error)

    // List retrieves list of Dashboards for a given namespace and list options.
    ListDashboard(ctx context.Context, opts ...client.ListOption) (*DashboardList, error)
}

// DashboardTransitionFunction instructs the DashboardWriter how to transition between an existing
// Dashboard object and a desired on an Upsert
type DashboardTransitionFunction func(existing, desired *Dashboard) error

// Writer knows how to create, delete, and update Dashboards.
type DashboardWriter interface {
    // Create saves the Dashboard object.
    CreateDashboard(ctx context.Context, obj *Dashboard, opts ...client.CreateOption) error

    // Delete deletes the Dashboard object.
    DeleteDashboard(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

    // Update updates the given Dashboard object.
    UpdateDashboard(ctx context.Context, obj *Dashboard, opts ...client.UpdateOption) error

    // Patch patches the given Dashboard object.
    PatchDashboard(ctx context.Context, obj *Dashboard, patch client.Patch, opts ...client.PatchOption) error

    // DeleteAllOf deletes all Dashboard objects matching the given options.
    DeleteAllOfDashboard(ctx context.Context, opts ...client.DeleteAllOfOption) error

    // Create or Update the Dashboard object.
    UpsertDashboard(ctx context.Context, obj *Dashboard, transitionFuncs ...DashboardTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Dashboard object.
type DashboardStatusWriter interface {
    // Update updates the fields corresponding to the status subresource for the
    // given Dashboard object.
    UpdateDashboardStatus(ctx context.Context, obj *Dashboard, opts ...client.UpdateOption) error

    // Patch patches the given Dashboard object's subresource.
    PatchDashboardStatus(ctx context.Context, obj *Dashboard, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Dashboards.
type DashboardClient interface {
    DashboardReader
    DashboardWriter
    DashboardStatusWriter
}

type dashboardClient struct {
    client client.Client
}

func NewDashboardClient(client client.Client) *dashboardClient {
    return &dashboardClient{client: client}
}


func (c *dashboardClient) GetDashboard(ctx context.Context, key client.ObjectKey) (*Dashboard, error) {
    obj := &Dashboard{}
    if err := c.client.Get(ctx, key, obj); err != nil {
        return nil, err
    }
    return obj, nil
}

func (c *dashboardClient) ListDashboard(ctx context.Context, opts ...client.ListOption) (*DashboardList, error) {
    list := &DashboardList{}
    if err := c.client.List(ctx, list, opts...); err != nil {
        return nil, err
    }
    return list, nil
}

func (c *dashboardClient) CreateDashboard(ctx context.Context, obj *Dashboard, opts ...client.CreateOption) error {
    return c.client.Create(ctx, obj, opts...)
}


func (c *dashboardClient) DeleteDashboard(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
    obj := &Dashboard{}
    obj.SetName(key.Name)
    obj.SetNamespace(key.Namespace)
    return c.client.Delete(ctx, obj, opts...)
}

func (c *dashboardClient) UpdateDashboard(ctx context.Context, obj *Dashboard, opts ...client.UpdateOption) error {
    return c.client.Update(ctx, obj, opts...)
}

func (c *dashboardClient) PatchDashboard(ctx context.Context, obj *Dashboard, patch client.Patch, opts ...client.PatchOption) error {
    return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *dashboardClient) DeleteAllOfDashboard(ctx context.Context, opts ...client.DeleteAllOfOption) error {
    obj := &Dashboard{}
    return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *dashboardClient) UpsertDashboard(ctx context.Context, obj *Dashboard, transitionFuncs ...DashboardTransitionFunction) error {
    genericTxFunc := func(existing, desired runtime.Object) error {
        for _, txFunc := range transitionFuncs {
            if err := txFunc(existing.(*Dashboard), desired.(*Dashboard)); err != nil {
                return err
            }
        }
        return nil
    }
    _, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
    return err
}

func (c *dashboardClient) UpdateDashboardStatus(ctx context.Context, obj *Dashboard, opts ...client.UpdateOption) error {
    return c.client.Status().Update(ctx, obj, opts...)
}

func (c *dashboardClient) PatchDashboardStatus(ctx context.Context, obj *Dashboard, patch client.Patch, opts ...client.PatchOption) error {
    return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides DashboardClients for multiple clusters.
type MulticlusterDashboardClient interface {
    // Cluster returns a DashboardClient for the given cluster
    Cluster(cluster string) (DashboardClient, error)
}

type multiclusterDashboardClient struct {
    client multicluster.Client
}

func NewMulticlusterDashboardClient(client multicluster.Client) MulticlusterDashboardClient {
    return &multiclusterDashboardClient{client: client}
}

func (m *multiclusterDashboardClient) Cluster(cluster string) (DashboardClient, error) {
    client, err := m.client.Cluster(cluster)
    if err != nil {
        return nil, err
    }
    return  NewDashboardClient(client), nil
}