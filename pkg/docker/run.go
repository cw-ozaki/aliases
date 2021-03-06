package docker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/k-kinzal/aliases/pkg/posix"
)

func quote(s string) string {
	length := len(s)
	if length >= 2 && (s[0] == '"' || s[0] == '\'') {
		if s[length-1] == '"' || s[length-1] == '\'' {
			s = s[1 : length-1]
		}
	}
	return strconv.Quote(s)
}

// RunOption is an option for executing docker run.
type RunOption struct {
	AddHost             []string
	Attach              []string
	BlkioWeight         *string
	BlkioWeightDevice   []string
	CIDFile             *string
	CPUPeriod           *string
	CPUQuota            *string
	CPURtPeriod         *string
	CPURtRuntime        *string
	CPUShares           *string
	CPUs                *string
	CPUsetCPUs          *string
	CPUsetMems          *string
	CapAdd              []string
	CapDrop             []string
	CgroupParent        *string
	DNS                 []string
	DNSOption           []string
	DNSSearch           []string
	Detach              *string
	DetachKeys          *string
	Device              []string
	DeviceCgroupRule    []string
	DeviceReadBPS       []string
	DeviceReadIOPS      []string
	DeviceWriteBPS      []string
	DeviceWriteIOPS     []string
	DisableContentTrust *string
	Domainname          *string
	Entrypoint          *string
	Env                 map[string]string
	EnvFile             []string
	Expose              []string
	GroupAdd            []string
	HealthCmd           *string
	HealthInterval      *string
	HealthRetries       *string
	HealthStartPeriod   *string
	HealthTimeout       *string
	Hostname            *string
	IP                  *string
	IP6                 *string
	IPC                 *string
	Init                *string
	Interactive         *string
	Isolation           *string
	KernelMemory        *string
	Label               map[string]string
	LabelFile           []string
	Link                []string
	LinkLocalIP         []string
	LogDriver           *string
	LogOpt              map[string]string
	MacAddress          *string
	Memory              *string
	MemoryReservation   *string
	MemorySwap          *string
	MemorySwappiness    *string
	Mount               map[string]string
	Name                *string
	Network             *string
	NetworkAlias        []string
	NoHealthcheck       *string
	OOMKillDisable      *string
	OOMScoreAdj         *string
	PID                 *string
	PidsLimit           *string
	Platform            *string
	Privileged          *string
	Publish             []string
	PublishAll          *string
	ReadOnly            *string
	Restart             *string
	Rm                  *string
	Runtime             *string
	SecurityOpt         map[string]string
	ShmSize             *string
	SigProxy            *string
	StopSignal          *string
	StopTimeout         *string
	StorageOpt          map[string]string
	Sysctl              map[string]string
	TTY                 *string
	Tmpfs               []string
	UTS                 *string
	Ulimit              map[string]string
	User                *string
	Userns              *string
	Volume              []string
	VolumeDriver        *string
	VolumesFrom         []string
	Workdir             *string
}

// Run returns a command that can execute docker run.
func (client *Client) Run(image string, args []string, option RunOption) *posix.Cmd {
	cmd := posix.Command(client.path, "run")

	for _, v := range option.AddHost {
		cmd.Args = append(cmd.Args, "--add-host", quote(v))
	}
	for _, v := range option.Attach {
		cmd.Args = append(cmd.Args, "--attach", quote(v))
	}
	if v := option.BlkioWeight; v != nil {
		cmd.Args = append(cmd.Args, "--blkio-weight", quote(*v))
	}
	for _, v := range option.BlkioWeightDevice {
		cmd.Args = append(cmd.Args, "--blkio-weight-device", quote(v))
	}
	for _, v := range option.CapAdd {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cap-add=%s", quote(v)))
	}
	for _, v := range option.CapDrop {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cap-drop=%s", quote(v)))
	}
	if v := option.CgroupParent; v != nil {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cgroup-parent==%s", quote(*v)))
	}
	if v := option.CIDFile; v != nil {
		cmd.Args = append(cmd.Args, "--cidfile", quote(*v))
	}
	if v := option.CPUPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-period", quote(*v))
	}
	if v := option.CPUQuota; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-quota", quote(*v))
	}
	if v := option.CPURtPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-rt-period", quote(*v))
	}
	if v := option.CPURtRuntime; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-rt-runtime", quote(*v))
	}
	if v := option.CPUShares; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-shares", quote(*v))
	}
	if v := option.CPUs; v != nil {
		cmd.Args = append(cmd.Args, "--cpus", quote(*v))
	}
	if v := option.CPUsetCPUs; v != nil {
		cmd.Args = append(cmd.Args, "--cpuset-cpus", quote(*v))
	}
	if v := option.CPUsetMems; v != nil {
		cmd.Args = append(cmd.Args, "--cpuset-mems", quote(*v))
	}
	if v := option.Detach; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--detach")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--detach\")", *v))
		}
	}
	if v := option.DetachKeys; v != nil {
		cmd.Args = append(cmd.Args, "--detach-keys", quote(*v))
	}
	for _, v := range option.Device {
		cmd.Args = append(cmd.Args, "--device", quote(v))
	}
	for _, v := range option.DeviceCgroupRule {
		cmd.Args = append(cmd.Args, "--device-cgroup-rule", quote(v))
	}
	for _, v := range option.DeviceReadBPS {
		cmd.Args = append(cmd.Args, "--device-read-bps", quote(v))
	}
	for _, v := range option.DeviceReadIOPS {
		cmd.Args = append(cmd.Args, "--device-read-iops", quote(v))
	}
	for _, v := range option.DeviceWriteBPS {
		cmd.Args = append(cmd.Args, "--device-write-bps", quote(v))
	}
	for _, v := range option.DeviceWriteIOPS {
		cmd.Args = append(cmd.Args, "--device-write-iops", quote(v))
	}
	if v := option.DisableContentTrust; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--disable-content-trust")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--disable-content-trust\")", *v))
		}
	}
	for _, v := range option.DNS {
		cmd.Args = append(cmd.Args, "--dns", quote(v))
	}
	for _, v := range option.DNSOption {
		cmd.Args = append(cmd.Args, "--dns-option", quote(v))
	}
	for _, v := range option.DNSSearch {
		cmd.Args = append(cmd.Args, "--dns-search", quote(v))
	}
	if v := option.Entrypoint; v != nil {
		cmd.Args = append(cmd.Args, "--entrypoint", quote(*v))
	}
	for _, k := range util.StringKeys(option.Env) {
		cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("%s=%s", k, quote(option.Env[k])))
	}
	for _, v := range option.EnvFile {
		cmd.Args = append(cmd.Args, "--env-file", quote(v))
	}
	for _, v := range option.Expose {
		cmd.Args = append(cmd.Args, "--expose", quote(v))
	}
	for _, v := range option.GroupAdd {
		cmd.Args = append(cmd.Args, "--group-add", quote(v))
	}
	if v := option.HealthCmd; v != nil {
		cmd.Args = append(cmd.Args, "--health-cmd", quote(*v))
	}
	if v := option.HealthInterval; v != nil {
		cmd.Args = append(cmd.Args, "--health-interval", quote(*v))
	}
	if v := option.HealthRetries; v != nil {
		cmd.Args = append(cmd.Args, "--health-retries", quote(*v))
	}
	if v := option.HealthStartPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--health-start-period", quote(*v))
	}
	if v := option.HealthTimeout; v != nil {
		cmd.Args = append(cmd.Args, "--health-timeout", quote(*v))
	}
	if v := option.Hostname; v != nil {
		cmd.Args = append(cmd.Args, "--hostname", quote(*v))
	}
	if v := option.Init; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--init")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--init\")", *v))
		}
	}
	if v := option.Interactive; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--interactive")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--interactive\")", *v))
		}
	}
	if v := option.IP; v != nil {
		cmd.Args = append(cmd.Args, "--ip", quote(*v))
	}
	if v := option.IP; v != nil {
		cmd.Args = append(cmd.Args, "--ip6", quote(*v))
	}
	if v := option.IPC; v != nil {
		cmd.Args = append(cmd.Args, "--ipc", quote(*v))
	}
	if v := option.Isolation; v != nil {
		cmd.Args = append(cmd.Args, "--isolation", quote(*v))
	}
	if v := option.KernelMemory; v != nil {
		cmd.Args = append(cmd.Args, "--kernel-memory", quote(*v))
	}
	for _, k := range util.StringKeys(option.Label) {
		cmd.Args = append(cmd.Args, "--label", fmt.Sprintf("%s=%s", k, quote(option.Label[k])))
	}
	for _, v := range option.LabelFile {
		cmd.Args = append(cmd.Args, "--label-file", quote(v))
	}
	for _, v := range option.Link {
		cmd.Args = append(cmd.Args, "--link", quote(v))
	}
	for _, v := range option.LinkLocalIP {
		cmd.Args = append(cmd.Args, "--link-loal-ip", quote(v))
	}
	if v := option.LogDriver; v != nil {
		cmd.Args = append(cmd.Args, "--log-driver", quote(*v))
	}
	for _, k := range util.StringKeys(option.LogOpt) {
		cmd.Args = append(cmd.Args, "--log-opt", fmt.Sprintf("%s=%s", k, quote(option.LogOpt[k])))
	}
	if v := option.MacAddress; v != nil {
		cmd.Args = append(cmd.Args, "--mac-address", quote(*v))
	}
	if v := option.Memory; v != nil {
		cmd.Args = append(cmd.Args, "--memory", quote(*v))
	}
	if v := option.MemoryReservation; v != nil {
		cmd.Args = append(cmd.Args, "--memory-reservation", quote(*v))
	}
	if v := option.MemorySwap; v != nil {
		cmd.Args = append(cmd.Args, "--memory-swap", quote(*v))
	}
	if v := option.MemorySwappiness; v != nil {
		cmd.Args = append(cmd.Args, "--memory-swappiness", quote(*v))
	}
	for _, k := range util.StringKeys(option.Mount) {
		cmd.Args = append(cmd.Args, "--mount", fmt.Sprintf("%s=%s", k, quote(option.Mount[k])))
	}
	if v := option.Name; v != nil {
		cmd.Args = append(cmd.Args, "--name", quote(*v))
	}
	if v := option.Network; v != nil {
		cmd.Args = append(cmd.Args, "--network", quote(*v))
	}
	for _, v := range option.NetworkAlias {
		cmd.Args = append(cmd.Args, "--network-alias", quote(v))
	}
	if v := option.NoHealthcheck; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--no-healthcheck")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--no-healthcheck\")", *v))
		}
	}
	if v := option.OOMKillDisable; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--oom-kill-disable")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--oom-kill-disable\")", *v))
		}
	}
	if v := option.OOMScoreAdj; v != nil {
		cmd.Args = append(cmd.Args, "--oom-secore-adj", quote(*v))
	}
	if v := option.PID; v != nil {
		cmd.Args = append(cmd.Args, "--pid", quote(*v))
	}
	if v := option.PidsLimit; v != nil {
		cmd.Args = append(cmd.Args, "--pids-limit", quote(*v))
	}
	if v := option.Platform; v != nil {
		cmd.Args = append(cmd.Args, "--platform", quote(*v))
	}
	if v := option.Privileged; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--privileged")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--privileged\")", *v))
		}
	}
	for _, v := range option.Publish {
		cmd.Args = append(cmd.Args, "--publish", quote(v))
	}
	if v := option.PublishAll; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--publish-all")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--publish-all\")", *v))
		}
	}
	if v := option.ReadOnly; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--readonly")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--readonly\")", *v))
		}
	}
	if v := option.Restart; v != nil {
		cmd.Args = append(cmd.Args, "--restart", quote(*v))
	}
	if v := option.Rm; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--rm")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--rm\")", *v))
		}
	}
	if v := option.Runtime; v != nil {
		cmd.Args = append(cmd.Args, "--runtime", quote(*v))
	}
	for _, k := range util.StringKeys(option.SecurityOpt) {
		cmd.Args = append(cmd.Args, "--security-opt", fmt.Sprintf("%s=%s", k, quote(option.SecurityOpt[k])))
	}
	if v := option.ShmSize; v != nil {
		cmd.Args = append(cmd.Args, "--shm-size", quote(*v))
	}
	if v := option.SigProxy; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--sig-proxy")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--sig-proxy\")", *v))
		}
	}
	if v := option.StopSignal; v != nil {
		cmd.Args = append(cmd.Args, "--stop-signal", quote(*v))
	}
	if v := option.StopTimeout; v != nil {
		cmd.Args = append(cmd.Args, "--stop-timeout", quote(*v))
	}
	for _, k := range util.StringKeys(option.StorageOpt) {
		cmd.Args = append(cmd.Args, "--storage-opt", fmt.Sprintf("%s=%s", k, quote(option.StorageOpt[k])))
	}
	for _, k := range util.StringKeys(option.Sysctl) {
		cmd.Args = append(cmd.Args, "--sysctl", fmt.Sprintf("%s=%s", k, quote(option.Sysctl[k])))
	}
	for _, v := range option.Tmpfs {
		cmd.Args = append(cmd.Args, "--tmpfs", quote(v))
	}
	if v := option.TTY; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--tty")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(%s && echo \"--tty\")", *v))
		}
	}
	for _, k := range util.StringKeys(option.Ulimit) {
		cmd.Args = append(cmd.Args, "--ulimit", fmt.Sprintf("%s=%s", k, quote(option.Ulimit[k])))
	}
	if v := option.User; v != nil {
		cmd.Args = append(cmd.Args, "--user", quote(*v))
	}
	if v := option.Userns; v != nil {
		cmd.Args = append(cmd.Args, "--userns", quote(*v))
	}
	if v := option.UTS; v != nil {
		cmd.Args = append(cmd.Args, "--uts", quote(*v))
	}
	for _, v := range option.Volume {
		cmd.Args = append(cmd.Args, "--volume", quote(v))
	}
	if v := option.VolumeDriver; v != nil {
		cmd.Args = append(cmd.Args, "--volume-driver", quote(*v))
	}
	for _, v := range option.VolumesFrom {
		cmd.Args = append(cmd.Args, "--volumes-from", quote(v))
	}
	if v := option.Workdir; v != nil {
		cmd.Args = append(cmd.Args, "--workdir", quote(*v))
	}
	for _, v := range option.VolumesFrom {
		cmd.Args = append(cmd.Args, "--volumes-from", quote(v))
	}
	cmd.Args = append(cmd.Args, image)
	for _, v := range args {
		if strings.HasPrefix(v, "'") && strings.HasSuffix(v, "'") {
			cmd.Args = append(cmd.Args, v)
		} else if strings.Contains(v, " ") {
			cmd.Args = append(cmd.Args, quote(v))
		} else {
			cmd.Args = append(cmd.Args, v)
		}
	}

	return cmd
}
