
---
title: "Gloo Mesh"
description: Reference for Helm values.
weight: 2
---

|Option|Type|Default Value|Description|
|------|----|-----------|-------------|
|glooMeshOperatorArgs|struct|{"settingsRef":{"name":"settings","namespace":"gloo-mesh"}}|Command line argument to Gloo Mesh deployments.|
|glooMeshOperatorArgs.settingsRef|struct|{"name":"settings","namespace":"gloo-mesh"}|Name/namespace of the Settings object.|
|glooMeshOperatorArgs.settingsRef.name|string|settings|Name of the Settings object.|
|glooMeshOperatorArgs.settingsRef.namespace|string|gloo-mesh|Namespace of the Settings object.|
|settings|struct|{"mtls":{"istio":{"tlsMode":"ISTIO_MUTUAL"}},"networkingExtensionServers":[],"discovery":{"istio":{"ingressGatewayDetectors":{}}},"relay":{"enabled":false,"server":{"address":"","insecure":false,"reconnectOnNetworkFailures":false}}}|Values for the Settings object. See the [Settings API doc](../../../../api/github.com.solo-io.gloo-mesh.api.settings.v1.settings) for details.|
|settings.mtls|struct|{"istio":{"tls_mode":2}}||
|settings.mtls.istio|struct|{"tls_mode":2}||
|settings.mtls.istio.tls_mode|int32|2||
|settings.networking_extension_servers[]|[]ptr|null||
|settings.networking_extension_servers[]|struct| ||
|settings.networking_extension_servers[].address|string| ||
|settings.networking_extension_servers[].insecure|bool| ||
|settings.networking_extension_servers[].reconnect_on_network_failures|bool| ||
|settings.discovery|struct|{"istio":{}}||
|settings.discovery.istio|struct|{}||
|settings.discovery.istio.ingress_gateway_detectors|map[string, struct]| ||
|settings.discovery.istio.ingress_gateway_detectors.<MAP_KEY>|struct| ||
|settings.discovery.istio.ingress_gateway_detectors.<MAP_KEY>.gateway_workload_labels|map[string, string]| ||
|settings.discovery.istio.ingress_gateway_detectors.<MAP_KEY>.gateway_workload_labels.<MAP_KEY>|string| ||
|settings.discovery.istio.ingress_gateway_detectors.<MAP_KEY>.gateway_tls_port_name|string| ||
|settings.relay|struct|{"server":{}}||
|settings.relay.enabled|bool|false||
|settings.relay.server|struct|{}||
|settings.relay.server.address|string| ||
|settings.relay.server.insecure|bool|false||
|settings.relay.server.reconnect_on_network_failures|bool|false||
|disallowIntersectingConfig|bool|false|If true, Gloo Mesh will detect and report errors when outputting service mesh configuration that overlaps with existing config not managed by Gloo Mesh.|
|watchOutputTypes|bool|true|If true, Gloo Mesh will watch service mesh config types output by Gloo Mesh, and resync upon changes.|
|defaultMetricsPort|uint32|9091|The port on which to serve internal Prometheus metrics for the Gloo Mesh application. Set to 0 to disable.|
|verbose|bool|false|If true, enables verbose/debug logging.|
|discovery|struct|{"image":{"repository":"gloo-mesh","registry":"gcr.io/gloo-mesh","pullPolicy":"IfNotPresent"},"env":[{"name":"POD_NAMESPACE","valueFrom":{"fieldRef":{"fieldPath":"metadata.namespace"}}}],"resources":{"requests":{"cpu":"125m","memory":"256Mi"}},"sidecars":{},"serviceType":"ClusterIP","ports":{"metrics":9091},"customPodAnnotations":{"sidecar.istio.io/inject":"\"false\""}}|Configuration for the discovery deployment.|
|discovery|struct|{"image":{"repository":"gloo-mesh","registry":"gcr.io/gloo-mesh","pullPolicy":"IfNotPresent"},"env":[{"name":"POD_NAMESPACE","valueFrom":{"fieldRef":{"fieldPath":"metadata.namespace"}}}],"resources":{"requests":{"cpu":"125m","memory":"256Mi"}}}||
|discovery.-[]|[]string|["discovery","--metrics-port={{ $.Values.discovery.ports.metrics | default $.Values.defaultMetricsPort }}","--settings-name={{ $.Values.glooMeshOperatorArgs.settingsRef.name }}","--settings-namespace={{ $.Values.glooMeshOperatorArgs.settingsRef.namespace }}","--verbose={{ $.Values.verbose }}"]||
|discovery.-[]|string| ||
|discovery.-[]|[]struct|null||
|discovery.-[]|struct| ||
|discovery.-[].name|string| ||
|discovery.-[].readOnly|bool| ||
|discovery.-[].mountPath|string| ||
|discovery.-[].subPath|string| ||
|discovery.-[].mountPropagation|string| ||
|discovery.-[].subPathExpr|string| ||
|discovery.image|struct|{"repository":"gloo-mesh","registry":"gcr.io/gloo-mesh","pullPolicy":"IfNotPresent"}|Specify the container image|
|discovery.image.tag|string| |Tag for the container.|
|discovery.image.repository|string|gloo-mesh|Image name (repository).|
|discovery.image.registry|string|gcr.io/gloo-mesh|Image registry.|
|discovery.image.pullPolicy|string|IfNotPresent|Image pull policy.|
|discovery.image.pullSecret|string| |Image pull secret.|
|discovery.Env[]|slice|[{"name":"POD_NAMESPACE","valueFrom":{"fieldRef":{"fieldPath":"metadata.namespace"}}}]|Specify environment variables for the container. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#envvarsource-v1-core) for specification details.|
|discovery.resources|struct|{"requests":{"cpu":"125m","memory":"256Mi"}}|Specify container resource requirements. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#resourcerequirements-v1-core) for specification details.|
|discovery.resources.limits|map[string, struct]| ||
|discovery.resources.limits.<MAP_KEY>|struct| ||
|discovery.resources.limits.<MAP_KEY>|string| ||
|discovery.resources.requests|map[string, struct]| ||
|discovery.resources.requests.<MAP_KEY>|struct| ||
|discovery.resources.requests.<MAP_KEY>|string| ||
|discovery.resources.requests.cpu|struct|"125m"||
|discovery.resources.requests.cpu|string|DecimalSI||
|discovery.resources.requests.memory|struct|"256Mi"||
|discovery.resources.requests.memory|string|BinarySI||
|discovery.sidecars|map[string, struct]| |Configuration for the deployed containers.|
|discovery.sidecars.<MAP_KEY>|struct| |Configuration for the deployed containers.|
|discovery.sidecars.<MAP_KEY>.-[]|[]string| ||
|discovery.sidecars.<MAP_KEY>.-[]|string| ||
|discovery.sidecars.<MAP_KEY>.-[]|[]struct| ||
|discovery.sidecars.<MAP_KEY>.-[]|struct| ||
|discovery.sidecars.<MAP_KEY>.-[].name|string| ||
|discovery.sidecars.<MAP_KEY>.-[].readOnly|bool| ||
|discovery.sidecars.<MAP_KEY>.-[].mountPath|string| ||
|discovery.sidecars.<MAP_KEY>.-[].subPath|string| ||
|discovery.sidecars.<MAP_KEY>.-[].mountPropagation|string| ||
|discovery.sidecars.<MAP_KEY>.-[].subPathExpr|string| ||
|discovery.sidecars.<MAP_KEY>.image|struct| |Specify the container image|
|discovery.sidecars.<MAP_KEY>.image.tag|string| |Tag for the container.|
|discovery.sidecars.<MAP_KEY>.image.repository|string| |Image name (repository).|
|discovery.sidecars.<MAP_KEY>.image.registry|string| |Image registry.|
|discovery.sidecars.<MAP_KEY>.image.pullPolicy|string| |Image pull policy.|
|discovery.sidecars.<MAP_KEY>.image.pullSecret|string| |Image pull secret.|
|discovery.sidecars.<MAP_KEY>.Env[]|slice| |Specify environment variables for the container. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#envvarsource-v1-core) for specification details.|
|discovery.sidecars.<MAP_KEY>.resources|struct| |Specify container resource requirements. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#resourcerequirements-v1-core) for specification details.|
|discovery.sidecars.<MAP_KEY>.resources.limits|map[string, struct]| ||
|discovery.sidecars.<MAP_KEY>.resources.limits.<MAP_KEY>|struct| ||
|discovery.sidecars.<MAP_KEY>.resources.limits.<MAP_KEY>|string| ||
|discovery.sidecars.<MAP_KEY>.resources.requests|map[string, struct]| ||
|discovery.sidecars.<MAP_KEY>.resources.requests.<MAP_KEY>|struct| ||
|discovery.sidecars.<MAP_KEY>.resources.requests.<MAP_KEY>|string| ||
|discovery.serviceType|string|ClusterIP|Specify the service type. Can be either "ClusterIP", "NodePort", "LoadBalancer", or "ExternalName".|
|discovery.ports|map[string, uint32]| |Specify service ports as a map from port name to port number.|
|discovery.ports.<MAP_KEY>|uint32| |Specify service ports as a map from port name to port number.|
|discovery.ports.metrics|uint32|9091|Specify service ports as a map from port name to port number.|
|discovery.customPodLabels|map[string, string]| |Custom labels for the pod|
|discovery.customPodLabels.<MAP_KEY>|string| |Custom labels for the pod|
|discovery.customPodAnnotations|map[string, string]| |Custom annotations for the pod|
|discovery.customPodAnnotations.<MAP_KEY>|string| |Custom annotations for the pod|
|discovery.customPodAnnotations.sidecar.istio.io/inject|string|"false"|Custom annotations for the pod|
|discovery.customDeploymentLabels|map[string, string]| |Custom labels for the deployment|
|discovery.customDeploymentLabels.<MAP_KEY>|string| |Custom labels for the deployment|
|discovery.customDeploymentAnnotations|map[string, string]| |Custom annotations for the deployment|
|discovery.customDeploymentAnnotations.<MAP_KEY>|string| |Custom annotations for the deployment|
|discovery.customServiceLabels|map[string, string]| |Custom labels for the service|
|discovery.customServiceLabels.<MAP_KEY>|string| |Custom labels for the service|
|discovery.customServiceAnnotations|map[string, string]| |Custom annotations for the service|
|discovery.customServiceAnnotations.<MAP_KEY>|string| |Custom annotations for the service|
|networking|struct|{"image":{"repository":"gloo-mesh","registry":"gcr.io/gloo-mesh","pullPolicy":"IfNotPresent"},"env":[{"name":"POD_NAMESPACE","valueFrom":{"fieldRef":{"fieldPath":"metadata.namespace"}}}],"resources":{"requests":{"cpu":"125m","memory":"256Mi"}},"sidecars":{},"serviceType":"","ports":{},"customPodAnnotations":{"sidecar.istio.io/inject":"\"false\""}}|Configuration for the networking deployment.|
|networking|struct|{"image":{"repository":"gloo-mesh","registry":"gcr.io/gloo-mesh","pullPolicy":"IfNotPresent"},"env":[{"name":"POD_NAMESPACE","valueFrom":{"fieldRef":{"fieldPath":"metadata.namespace"}}}],"resources":{"requests":{"cpu":"125m","memory":"256Mi"}}}||
|networking.-[]|[]string|["networking","--metrics-port={{ $.Values.defaultMetricsPort }}","--settings-name={{ $.Values.glooMeshOperatorArgs.settingsRef.name }}","--settings-namespace={{ $.Values.glooMeshOperatorArgs.settingsRef.namespace }}","--verbose={{ $.Values.verbose }}","--disallow-intersecting-config={{ $.Values.disallowIntersectingConfig }}","--watch-output-types={{ $.Values.watchOutputTypes }}"]||
|networking.-[]|string| ||
|networking.-[]|[]struct|null||
|networking.-[]|struct| ||
|networking.-[].name|string| ||
|networking.-[].readOnly|bool| ||
|networking.-[].mountPath|string| ||
|networking.-[].subPath|string| ||
|networking.-[].mountPropagation|string| ||
|networking.-[].subPathExpr|string| ||
|networking.image|struct|{"repository":"gloo-mesh","registry":"gcr.io/gloo-mesh","pullPolicy":"IfNotPresent"}|Specify the container image|
|networking.image.tag|string| |Tag for the container.|
|networking.image.repository|string|gloo-mesh|Image name (repository).|
|networking.image.registry|string|gcr.io/gloo-mesh|Image registry.|
|networking.image.pullPolicy|string|IfNotPresent|Image pull policy.|
|networking.image.pullSecret|string| |Image pull secret.|
|networking.Env[]|slice|[{"name":"POD_NAMESPACE","valueFrom":{"fieldRef":{"fieldPath":"metadata.namespace"}}}]|Specify environment variables for the container. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#envvarsource-v1-core) for specification details.|
|networking.resources|struct|{"requests":{"cpu":"125m","memory":"256Mi"}}|Specify container resource requirements. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#resourcerequirements-v1-core) for specification details.|
|networking.resources.limits|map[string, struct]| ||
|networking.resources.limits.<MAP_KEY>|struct| ||
|networking.resources.limits.<MAP_KEY>|string| ||
|networking.resources.requests|map[string, struct]| ||
|networking.resources.requests.<MAP_KEY>|struct| ||
|networking.resources.requests.<MAP_KEY>|string| ||
|networking.resources.requests.cpu|struct|"125m"||
|networking.resources.requests.cpu|string|DecimalSI||
|networking.resources.requests.memory|struct|"256Mi"||
|networking.resources.requests.memory|string|BinarySI||
|networking.sidecars|map[string, struct]| |Configuration for the deployed containers.|
|networking.sidecars.<MAP_KEY>|struct| |Configuration for the deployed containers.|
|networking.sidecars.<MAP_KEY>.-[]|[]string| ||
|networking.sidecars.<MAP_KEY>.-[]|string| ||
|networking.sidecars.<MAP_KEY>.-[]|[]struct| ||
|networking.sidecars.<MAP_KEY>.-[]|struct| ||
|networking.sidecars.<MAP_KEY>.-[].name|string| ||
|networking.sidecars.<MAP_KEY>.-[].readOnly|bool| ||
|networking.sidecars.<MAP_KEY>.-[].mountPath|string| ||
|networking.sidecars.<MAP_KEY>.-[].subPath|string| ||
|networking.sidecars.<MAP_KEY>.-[].mountPropagation|string| ||
|networking.sidecars.<MAP_KEY>.-[].subPathExpr|string| ||
|networking.sidecars.<MAP_KEY>.image|struct| |Specify the container image|
|networking.sidecars.<MAP_KEY>.image.tag|string| |Tag for the container.|
|networking.sidecars.<MAP_KEY>.image.repository|string| |Image name (repository).|
|networking.sidecars.<MAP_KEY>.image.registry|string| |Image registry.|
|networking.sidecars.<MAP_KEY>.image.pullPolicy|string| |Image pull policy.|
|networking.sidecars.<MAP_KEY>.image.pullSecret|string| |Image pull secret.|
|networking.sidecars.<MAP_KEY>.Env[]|slice| |Specify environment variables for the container. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#envvarsource-v1-core) for specification details.|
|networking.sidecars.<MAP_KEY>.resources|struct| |Specify container resource requirements. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#resourcerequirements-v1-core) for specification details.|
|networking.sidecars.<MAP_KEY>.resources.limits|map[string, struct]| ||
|networking.sidecars.<MAP_KEY>.resources.limits.<MAP_KEY>|struct| ||
|networking.sidecars.<MAP_KEY>.resources.limits.<MAP_KEY>|string| ||
|networking.sidecars.<MAP_KEY>.resources.requests|map[string, struct]| ||
|networking.sidecars.<MAP_KEY>.resources.requests.<MAP_KEY>|struct| ||
|networking.sidecars.<MAP_KEY>.resources.requests.<MAP_KEY>|string| ||
|networking.serviceType|string| |Specify the service type. Can be either "ClusterIP", "NodePort", "LoadBalancer", or "ExternalName".|
|networking.ports|map[string, uint32]| |Specify service ports as a map from port name to port number.|
|networking.ports.<MAP_KEY>|uint32| |Specify service ports as a map from port name to port number.|
|networking.customPodLabels|map[string, string]| |Custom labels for the pod|
|networking.customPodLabels.<MAP_KEY>|string| |Custom labels for the pod|
|networking.customPodAnnotations|map[string, string]| |Custom annotations for the pod|
|networking.customPodAnnotations.<MAP_KEY>|string| |Custom annotations for the pod|
|networking.customPodAnnotations.sidecar.istio.io/inject|string|"false"|Custom annotations for the pod|
|networking.customDeploymentLabels|map[string, string]| |Custom labels for the deployment|
|networking.customDeploymentLabels.<MAP_KEY>|string| |Custom labels for the deployment|
|networking.customDeploymentAnnotations|map[string, string]| |Custom annotations for the deployment|
|networking.customDeploymentAnnotations.<MAP_KEY>|string| |Custom annotations for the deployment|
|networking.customServiceLabels|map[string, string]| |Custom labels for the service|
|networking.customServiceLabels.<MAP_KEY>|string| |Custom labels for the service|
|networking.customServiceAnnotations|map[string, string]| |Custom annotations for the service|
|networking.customServiceAnnotations.<MAP_KEY>|string| |Custom annotations for the service|
