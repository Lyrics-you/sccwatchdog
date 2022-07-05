-- all
# deployment name
kubectl get deployment -n "namespace" -o jsonpath="{.items[*].metadata.name}"
# container namespace
kubectl get deployment -n "namespace" -o jsonpath="{.items[*].metadata.namespace}"
# deployment containers image
kubectl get deployment -n "namespace" -o jsonpath="{.items[*].spec.template.spec.containers[*].image}"

-- specified deployment

# deployment containers image
kubectl get deployment -n "namespace" "deployment"  -o jsonpath="{.spec.template.spec.containers[*].image}"
# deployment containers name
kubectl get deployment -n "namespace" "deployment"  -o jsonpath="{.spec.template.spec.containers[*].name}"
# deployment status
kubectl get deployment -n "namespace" "deployment"  -o jsonpath="{.spec.template.status}"
# deployment lastUpdateTime
kubectl get deployment -n "namespace" "deployment"  -o jsonpath="{.status.conditions[-1].lastUpdateTime}"


-- jsonpath
#image,lastUpdateTime
{range.items[*]}[{.spec.template.spec.containers[*].image},{.status.conditions[-1].lastUpdateTime}]{end}
# name,image,lastUpdateTime
{range.items[*]}[{.metadata.name},{.spec.template.spec.containers[*].image},{.status.conditions[-1].lastUpdateTime}]{end}

-- set deployment 
# kubectl set image -n "namespace" deployment.apps/{deployment名称} {镜像名称}:={镜像名称}:{版本}
kubectl set image -n mergemultx deployments.apps/mergemultx1 mergemultx=loveyoutruehappy/mergemultx:v1.2

-- restart deployment
#基本思路就是给Container添加一个无关紧要的环境变量，这个环境变量的值就是时间戳，而这个时间戳则是每次执行上述命令的系统当前时间。
kubectl patch deployment <deployment-name> \
  -p '{"spec":{"template":{"spec":{"containers":[{"name":"<container-name>","env":[{"name":"SWD_RESTART","value":"'$(date +%s)'"}]}]}}}}'

mergemultx.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mergemultx1
  namespace : mergemultx
  labels:
    app: mergemultx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: mergemultx
  template:
    metadata:
      labels:
        app: mergemultx
    spec:
      containers:
      - name: mergemultx
        image: loveyoutruehappy/mergemultx:v1.1
        ports:
        - containerPort: 39789
---

[root@centos-7-03 k8s]# kubectl create namespace mergemultx
[root@centos-7-03 k8s]# kubectl get deployment -n mergemultx -o wide
NAME          READY   UP-TO-DATE   AVAILABLE   AGE     CONTAINERS   IMAGES                             SELECTOR
mergemultx1   3/3     3            3           3m6s    mergemultx   loveyoutruehappy/mergemultx:v1.5   app=mergemultx
mergemultx2   3/3     3            3           2m49s   mergemultx   loveyoutruehappy/mergemultx:v1.2   app=mergemultx
mergemultx3   3/3     3            3           2m23s   mergemultx   loveyoutruehappy/mergemultx:v1.3   app=mergemultx

kubectl patch deployment -n mergemultx mergemultx1 \
  -p '{"spec":{"template":{"spec":{"containers":[{"name":"mergemultx","env":[{"name":"RESTART_","value":"'$(date +%s)'"}]}]}}}}'

kubectl patch deployment -n mergemultx mergemultx1 -p '{"spec":{"template":{"spec":{"containers":[{"name":"mergemultx","env":[{"name":"SWD_RESTART","value":"'$(date +%s)'"}]}]}}}}'

kubectl patch deployment mergemultx1 -p '{"spec":{"template":{"spec":{"containers":[{"name":"mergemultx","env":[{"name":"SWD_RESTART","value":"'$(date +%s)'"}]}]}}}}'
