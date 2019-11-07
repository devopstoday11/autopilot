# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [autopilot.proto](#autopilot.proto)
    - [AutoPilotProject](#autopilot.AutoPilotProject)
    - [Parameter](#autopilot.Parameter)
    - [Phase](#autopilot.Phase)
  
  
  
  

- [autopilot-operator.proto](#autopilot-operator.proto)
    - [AutoPilotOperator](#autopilot.AutoPilotOperator)
  
    - [MeshProvider](#autopilot.MeshProvider)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="autopilot.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## autopilot.proto
autopilot-operator.proto defines the API Schema for the autopilot.yaml configuration file.
This file is used to generate and re-generate the project structure, as well
as execute tasks related to build and deployment. It can be consumed
both via the `ap` CLI as well as in `codegen` packages.


<a name="autopilot.AutoPilotProject"></a>

### AutoPilotProject
The AutoPilotProject file is the root configuration file for the project itself.

This file will be used to build and deploy the autopilot operator.
It is loaded automatically by the autopilot CLI. Its
default location is 'autopilot.yaml'


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | the name (kubernetes Kind) of the top-level CRD for the operator Specified via the `ap init <Kind>` command |
| apiVersion | [string](#string) |  | the ApiVersion of the top-level CRD for the operator |
| operatorName | [string](#string) |  | the name of the Operator this is used to name and label loggers, k8s resources, and metrics exposed by the operator. Should be [valid Kube resource names](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names). |
| phases | [Phase](#autopilot.Phase) | repeated | Each phase represents a different stage in the lifecycle of the CRD (e.g. Pending/Succeeded/Failed).

Each phase specifies a unique name and its own set of inputs and outputs. |
| enableFinalizer | [bool](#bool) |  | enable use of a Finalizer to handle object deletion |
| customParameters | [Parameter](#autopilot.Parameter) | repeated | custom Parameters which extend AutoPilot's builtin types |






<a name="autopilot.Parameter"></a>

### Parameter
Custom Parameters allow code to be generated
for inputs/outputs that are not built-in to AutoPilot.
These types must be Kubernetes-compatible Go structs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lowerName | [string](#string) |  | the fully lower-case name of this resource e.g. "pods", "services", "replicasets", "configmaps" |
| singleName | [string](#string) |  | the singular CamelCased name of the resource equivalent to Kind |
| pluralName | [string](#string) |  | the plural CamelCased name of the resource equivalent to the pluralized form of Kind |
| importPrefix | [string](#string) |  | import prefix used by generated code |
| package | [string](#string) |  | go package (import path) to the go struct for the resource |
| apiGroup | [string](#string) |  | Kubernetes API group for the resource e.g. "networking.istio.io" |
| isCrd | [bool](#bool) |  | indicates whether the resource is a CRD if true, the Resource will be added to the operator's runtime.Scheme |






<a name="autopilot.Phase"></a>

### Phase
MeshProviders provide an interface to monitoring and managing a specific
mesh.

AutoPilot does not abstract the mesh API - AutoPilot developers must
still reason able about Provider-specific CRDs. AutoPilot's job is to
abstract operational concerns such as discovering control plane configuration
and monitoring metrics.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name of the phase. must be unique |
| description | [string](#string) |  | description of the phase. used for comments and docs |
| initial | [bool](#bool) |  | indicates whether this is the initial phase of the system. exactly one phase must be the initial phase |
| inputs | [string](#string) | repeated | the set of inputs for this phase the inputs will be retrieved by the scheduler and passed to the worker as input parameters

custom inputs can be defined in the autopilot.yaml |
| outputs | [string](#string) | repeated | the set of outputs for this phase the inputs will be propagated to k8s storage (etcd) by the scheduler.

custom outputs can be defined in the autopilot.yaml |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="autopilot-operator.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## autopilot-operator.proto
autopilot-operator.proto defines the API Schema for the autopilot-operator.yaml configuration file.
this file provides the bootstrap configuration that is loaded to the
operator at boot-time/runtime


<a name="autopilot.AutoPilotOperator"></a>

### AutoPilotOperator
The AutoPilotOperator file is the bootstrap
Configuration file for the Operator.
It is stored and mounted to the operator as a Kubernetes ConfigMap.
The Operator will hot-reload when the configuration file changes.
Default name is 'autopilot-operator.yaml' and should be stored in the project root.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | version of the operator used for reporting, metrics, etc (can be any format) default is "0.0.1" |
| meshProvider | [MeshProvider](#autopilot.MeshProvider) |  | meshProvider determines how the operator will connect to a service mesh Default is "SMI" |
| workInterval | [google.protobuf.Duration](#google.protobuf.Duration) |  | workInterval to sets the interval at which CRD workers resync. Default is 5s |
| metricsAddr | [string](#string) |  | Serve metrics on this address. Set to empty string to disable metrics defaults to ":9090" |
| enableLeaderElection | [bool](#bool) |  | Enable leader election. This will prevent more than one operator from running at a time defaults to true |
| watchNamespace | [string](#string) |  | if non-empty, watchNamespace will restrict the Operator to watching resources in a single namespace if empty (default), the Operator must have Cluster-scope RBAC permissions (ClusterRole/Binding) |





 <!-- end messages -->


<a name="autopilot.MeshProvider"></a>

### MeshProvider
MeshProviders provide an interface to monitoring and managing a specific
mesh.
AutoPilot does not abstract the mesh API - AutoPilot developers must
still reason able about Provider-specific CRDs. AutoPilot's job is to
abstract operational concerns such as discovering control plane configuration
and monitoring metrics.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SMI | 0 | the Operator will utilize the Service Mesh Interface (SMI) for metrics and configuration. Compatible with multiple meshes (may require installation of an SMI Adapter). |
| Istio | 1 | the Operator will utilize Istio mesh for metrics and configuration |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |