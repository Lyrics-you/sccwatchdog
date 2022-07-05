# sccwatchdog

Sccwatchdog is a free and open source scc deployment watcher system,designed to watch scc deployment image version and updatetime with speed and efficiency.You can also restart deployment and set container image by it.

<img src=".\image-20220705143644713.png" alt="image-20220705143644713" style="zoom: 67%;" />

## 参数

`swd --help`

```shell
Sccwatchdog is a free and open source scc deployment watcher system,
designed to watch scc deployment image version and updatetime with speed and efficiency.
You can also restart deployment and set container image by it.

Usage:
  swd [flags]
  swd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  restart     Restart by namespace and deployment
  set         Set image by namespace and deployment
  show        Show information by namespace and deployments
  version     Version subcommand show swd version info
  watch       Watch changed and restarted by namespace and deployments 

Flags:
  -h, --help      help for swd
  -v, --version   sccwatchdog version

Use "swd [command] --help" for more information about a command.
```

## version

`swd version --help`

```shell
eg.: swd version [-d]

Usage:
  swd version [flags]

Flags:
  -d, --description   history description
  -h, --help          help for version

Global Flags:
  -v, --version   sccwatchdog version
```

```shell
[root@centos-7-03 sccwatchdog]# ./swd version
Name          : SCCWatchDog🐶
Version       : 0.4.0
Email         : Leyuan.Jia@Outlook.com
[root@centos-7-03 sccwatchdog]# ./swd version -d
Name          : SCCWatchDog🐶
Version       : 0.4.0
Description   : add restart deployment and set image function
```

## show

`swd show --help`

```shell
Show all deployments in namespace
eg.: swd show [-n <namespace>(default:"default")]
Show specified deployments "deploy1 deploy2" in namespace
eg.: swd show [-n <namespace>(default:"default")] -d "deploy1 deploy2"

Usage:
  swd show [flags]

Flags:
  -d, --depolyment string   scc depolyment
  -h, --help                help for show
  -n, --namespace string    scc namespace

Global Flags:
  -v, --version   sccwatchdog version

```

```shell
[root@centos-7-03 sccwatchdog]# ./swd show -n mergemultx -d "mergemultx1 mergemultx2"
[mergemultx/mergemultx1: (loveyoutruehappy/mergemultx:v1.3) @ 2022-07-05 09:42:30 CST]
[mergemultx/mergemultx2: (loveyoutruehappy/mergemultx:v1.2) @ 2022-07-05 09:42:31 CST]
```

## watch

`swd watch --help`

```shell
Watch all deployments in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] [-s t]
Watch specified deployments "deploy1 deploy2" in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] -d "deploy1 deploy2" [-s t]

Usage:
  swd watch [flags]

Flags:
  -d, --depolyment string   scc depolyment
  -h, --help                help for watch
  -n, --namespace string    scc namespace
  -s, --second int          times interval (second)

Global Flags:
  -v, --version   sccwatchdog version

```

```shell
[root@centos-7-03 sccwatchdog]# ./swd watch -n mergemultx -d "mergemultx1 mergemultx2" -s 5
[22-07-05 14:44:45.756][dog/watch.go:49 @WatchStart()][INFO]mergemultx1 has changed,loveyoutruehappy/mergemultx:v1.3==>loveyoutruehappy/mergemultx:v1.2 at 2022-07-05 14:44:43 CST
[22-07-05 14:44:50.885][dog/watch.go:49 @WatchStart()][INFO]mergemultx1 has restarted,loveyoutruehappy/mergemultx:v1.2 at 2022-07-05 14:44:48 CST,ReplicaSet "mergemultx1-675d478458" is progressing.
[22-07-05 14:45:36.483][dog/watch.go:49 @WatchStart()][INFO]mergemultx2 has restarted,loveyoutruehappy/mergemultx:v1.2 at 2022-07-05 14:45:31 CST,Deployment does not have minimum availability.
```

## set

`swd set --help`

```shell
eg.: swd set -n <namespace> -d <deployment> [-c <container>(default:first container,else <deployment>)] -i <image>

Usage:
  swd set [flags]

Flags:
  -c, --container string    depolyment container
  -d, --depolyment string   scc depolyment
  -h, --help                help for set
  -i, --image string        container image
  -n, --namespace string    scc namespace

Global Flags:
  -v, --version   sccwatchdog version

```

```shell
[root@centos-7-03 sccwatchdog]# ./swd set -n mergemultx -d mergemultx1 -i loveyoutruehappy/mergemultx:v1.3
[22-07-05 14:44:34.629][cmd/set.go:38 @setDeploymentImage()][INFO]mergemultx1(mergemultx) image not changed
[root@centos-7-03 sccwatchdog]# ./swd set -n mergemultx -d mergemultx1 -i loveyoutruehappy/mergemultx:v1.2
[22-07-05 14:44:43.028][cmd/set.go:38 @setDeploymentImage()][INFO]deployment.apps/mergemultx1 image updated
```

## restart

`swd restart --help`

```shell
eg.: swd set -n <namespace> -d <deployment> [-c <container>(default:first container,else <deployment>)] -i <image>

Usage:
  swd set [flags]

Flags:
  -c, --container string    depolyment container
  -d, --depolyment string   scc depolyment
  -h, --help                help for set
  -i, --image string        container image
  -n, --namespace string    scc namespace

Global Flags:
  -v, --version   sccwatchdog version

```

```shell
[root@centos-7-03 sccwatchdog]# ./swd restart -n mergemultx -d mergemultx2
[22-07-05 14:46:37.419][cmd/restart.go:37 @restartDeployment()][INFO]deployment.apps/mergemultx2 patched
```

