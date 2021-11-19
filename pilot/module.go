package pilot

type ENV struct{
	Key string
	Value string
}
type Mount struct{
	Container_path string
	Host_ptah string
	Readonly string
}
type InfoStru  struct{
	SandboxID string
	Pid   string
	Config struct {
		Envs [] ENV
		Mounts [] Mount
		Labels struct {
			ContainerName string `json:"io.kubernetes.container.name",omitempty`
			PodName string `json:"io.kubernetes.pod.name",omitempty`
			PodNamespace string `json:"io.kubernetes.pod.namespace",omitempty`
			PodUid string `json:"io.kubernetes.pod.uid",omitempty`
		}
	}
	RuntimeSpec struct{
		OciVersion string
		Process struct{
			Args [] string
			Env []  string
		}
		Mounts [] struct{
			Destination string
			Source string
			Type string
		}
	}

}
