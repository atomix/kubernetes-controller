package k8s

import (
	api "github.com/atomix/atomix-api/proto/atomix/controller"
	"github.com/atomix/atomix-k8s-controller/pkg/apis/k8s/v1alpha1"
	"github.com/atomix/atomix-k8s-controller/pkg/controller/protocol"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/proto"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetPartitionSetName(id *api.PartitionGroupId) string {
	return id.Name
}

func GetPartitionSetNamespace(id *api.PartitionGroupId) string {
	if id.Namespace != "" {
		return id.Namespace
	} else {
		return DefaultNamespace
	}
}

func GetPartitionSetNamespacedName(id *api.PartitionGroupId) types.NamespacedName {
	return types.NamespacedName{
		Name:      GetPartitionSetName(id),
		Namespace: GetPartitionSetNamespace(id),
	}
}

// GetPartitionSetServiceName returns the service name for a partition set
func GetPartitionSetServiceName(set *v1alpha1.PartitionSet) string {
	return set.Name
}

// GetPartitionSetServiceNamespacedName returns the namespaced service name for a partition set
func GetPartitionSetServiceNamespacedName(set *v1alpha1.PartitionSet) types.NamespacedName {
	return types.NamespacedName{
		Name:      GetPartitionSetServiceName(set),
		Namespace: set.Namespace,
	}
}

// NewPartitionSetService returns a new service for a partition set
func NewPartitionSetService(set *v1alpha1.PartitionSet) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      GetPartitionSetServiceName(set),
			Namespace: set.Namespace,
			Labels:    set.Labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "api",
					Port: 5678,
				},
			},
		},
	}
}

// NewPartitionSetEndpoints returns an Endpoints object for the given partition set
func NewPartitionSetEndpoints(set *v1alpha1.PartitionSet) *corev1.Endpoints {
	return &corev1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      set.Name,
			Namespace: set.Namespace,
		},
		Subsets: []corev1.EndpointSubset{},
	}
}

// NewPartitionSetEndpointPorts returns an EndpointPort for a partition set
func NewPartitionSetEndpointPorts() []corev1.EndpointPort {
	return []corev1.EndpointPort{
		{
			Name: "api",
			Port: 5678,
		},
	}
}

func NewPartitionProto(p *v1alpha1.Partition) (*api.Partition, error) {
	id, err := getPartitionIdFromAnnotation(p)
	if err != nil {
		return nil, err
	}
	return &api.Partition{
		PartitionID: int32(id),
		Endpoints: []*api.PartitionEndpoint{
			{
				Host: GetPartitionServiceName(p),
				Port: 5678,
			},
		},
	}, nil
}

func NewPartitionGroupProtoFromSet(set *v1alpha1.PartitionSet, protocols *protocol.ProtocolManager) (*api.PartitionGroup, error) {
	spec, err := newPartitionGroupSpecProto(set, protocols)
	if err != nil {
		return nil, err
	}
	return &api.PartitionGroup{
		ID: &api.PartitionGroupId{
			Name:      set.Name,
			Namespace: set.Namespace,
		},
		Spec: spec,
	}, nil
}

func newPartitionGroupSpecProto(set *v1alpha1.PartitionSet, protocols *protocol.ProtocolManager) (*api.PartitionGroupSpec, error) {
	protocol, err := protocols.GetProtocolByName(set.Spec.Template.Spec.Protocol)
	if err != nil {
		return nil, err
	}

	message, err := protocol.YamlToMessage([]byte(set.Spec.Template.Spec.Config))
	if err != nil {
		return nil, err
	}
	bytes, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &api.PartitionGroupSpec{
		Partitions:    uint32(set.Spec.Partitions),
		PartitionSize: uint32(set.Spec.Template.Spec.Size),
		Protocol: &gogotypes.Any{
			TypeUrl: "type.googleapis.com/" + proto.MessageName(message),
			Value:   bytes,
		},
	}, nil
}

func NewPartitionSetFromProto(id *api.PartitionGroupId, pbspec *api.PartitionGroupSpec, protocols *protocol.ProtocolManager) (*v1alpha1.PartitionSet, error) {
	protocol, err := protocols.GetProtocolByType(pbspec.Protocol.TypeUrl)
	if err != nil {
		return nil, err
	}

	ns := id.Namespace
	if ns == "" {
		ns = DefaultNamespace
	}

	yaml, err := protocol.ProtoToYaml(pbspec.Protocol.Value)
	if err != nil {
		return nil, err
	}

	return &v1alpha1.PartitionSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      id.Name,
			Namespace: id.Namespace,
		},
		Spec: v1alpha1.PartitionSetSpec{
			Partitions: int(pbspec.Partitions),
			Template: v1alpha1.PartitionTemplateSpec{
				Spec: v1alpha1.PartitionSpec{
					Size:     int32(pbspec.PartitionSize),
					Protocol: protocol.Name,
					Image:    protocol.Image,
					Config:   string(yaml),
				},
			},
		},
	}, nil
}