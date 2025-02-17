package federation_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	. "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/mesh/federation"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	skv1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	"github.com/solo-io/skv2/pkg/ezkube"
	networkingv1alpha3spec "istio.io/api/networking/v1alpha3"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("FederationTranslator", func() {
	ctx := context.TODO()

	It("translates federation resources for a VirtualMesh", func() {

		namespace := "namespace"
		clusterName := "cluster"

		mesh := &discoveryv1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "config-namespace",
				Name:      "federated-mesh",
			},
			Spec: discoveryv1.MeshSpec{
				Type: &discoveryv1.MeshSpec_Istio_{Istio: &discoveryv1.MeshSpec_Istio{
					SmartDnsProxyingEnabled: true,
					Installation: &discoveryv1.MeshInstallation{
						Namespace: namespace,
						Cluster:   clusterName,
						Version:   "1.8.1",
					},
					IngressGateways: []*discoveryv1.MeshSpec_Istio_IngressGatewayInfo{{
						ExternalAddressType: &discoveryv1.MeshSpec_Istio_IngressGatewayInfo_DnsName{
							DnsName: "mesh-gateway.dns.name",
						},
						ExternalTlsPort:  8181,
						TlsContainerPort: 9191,
						WorkloadLabels:   map[string]string{"gatewaylabels": "righthere"},
					}},
				}},
			},
		}

		clientMesh := &discoveryv1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "config-namespace",
				Name:      "client-mesh",
			},
			Spec: discoveryv1.MeshSpec{
				Type: &discoveryv1.MeshSpec_Istio_{Istio: &discoveryv1.MeshSpec_Istio{
					SmartDnsProxyingEnabled: true,
					Installation: &discoveryv1.MeshInstallation{
						Namespace: "remote-namespace",
						Cluster:   "remote-cluster",
					},
				}},
			},
		}

		meshRef := ezkube.MakeObjectRef(mesh)
		clientMeshRef := ezkube.MakeObjectRef(clientMesh)

		vMesh := &discoveryv1.MeshStatus_AppliedVirtualMesh{
			Ref: &skv2corev1.ObjectRef{
				Name:      "my-virtual-mesh",
				Namespace: "config-namespace",
			},
			Spec: &v1.VirtualMeshSpec{
				Meshes: []*skv2corev1.ObjectRef{
					meshRef,
					clientMeshRef,
				},
				Federation: &v1.VirtualMeshSpec_Federation{
					HostnameSuffix: "soloio",
				},
			},
		}

		kubeCluster := &skv1alpha1.KubernetesCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      clusterName,
				Namespace: defaults.GetPodNamespace(),
			},
		}

		in := input.NewInputLocalSnapshotManualBuilder("ignored").
			AddMeshes(discoveryv1.MeshSlice{mesh, clientMesh}).
			AddKubernetesClusters(skv1alpha1.KubernetesClusterSlice{kubeCluster}).
			Build()

		t := NewTranslator(ctx)

		outputs := istio.NewBuilder(context.TODO(), "")
		t.Translate(
			in,
			mesh,
			vMesh,
			outputs,
			nil, // no reports expected
		)

		Expect(outputs.GetGateways().Length()).To(Equal(1))
		Expect(outputs.GetGateways().List()[0]).To(Equal(expectedGateway))
	})
})

var expectedGateway = &networkingv1alpha3.Gateway{
	ObjectMeta: metav1.ObjectMeta{
		Name:        "my-virtual-mesh-config-namespace",
		Namespace:   "namespace",
		ClusterName: "cluster",
		Labels:      metautils.TranslatedObjectLabels(),
		Annotations: map[string]string{
			metautils.ParentLabelkey: `{"networking.mesh.gloo.solo.io/v1, Kind=VirtualMesh":[{"name":"my-virtual-mesh","namespace":"config-namespace"}]}`,
		},
	},
	Spec: networkingv1alpha3spec.Gateway{
		Servers: []*networkingv1alpha3spec.Server{
			{
				Port: &networkingv1alpha3spec.Port{
					Number:   9191,
					Protocol: "TLS",
					Name:     "tls",
				},
				Hosts: []string{
					"*.soloio",
				},
				Tls: &networkingv1alpha3spec.ServerTLSSettings{
					Mode: networkingv1alpha3spec.ServerTLSSettings_AUTO_PASSTHROUGH,
				},
			},
		},
		Selector: map[string]string{"gatewaylabels": "righthere"},
	},
}
