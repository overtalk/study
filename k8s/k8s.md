# Kubenetes

- metadata : 用于唯一区分对象的元数据，包括：name，UID和namespace
- labels：是一个个的key/value对，定义这样的label到Pod后，其他控制器对象可以通过这样的label来定位到此Pod，从而对Pod进行管理。（参见Deployment等控制器对象）
- spec： 其它描述信息，包含Pod中运行的容器，容器中运行的应用等等。不同类型的对象拥有不同的spec定义。详情参见API文档：https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.9/


## kubectl 命令
- 创建对象（json/yaml）
    - kubectl create -f xxx
- 显示和查找资源
    - kubectl get
        - kubectl get services                          # 列出所有 namespace 中的所有 service
        - kubectl get pods --all-namespaces             # 列出所有 namespace 中的所有 pod
        - kubectl get pods -o wide                      # 列出所有 pod 并显示详细信息
        - kubectl get deployment my-dep                 # 列出指定 deployment
        - kubectl get pods --include-uninitialized      # 列出该 namespace 中的所有 pod 包括未初始化的
- 删除资源
    - kubectl delete
        - kubectl delete -f ./pod.json                                              # 删除 pod.json 文件中定义的类型和名称的 pod
        - kubectl delete pod,service baz foo                                        # 删除名为“baz”的 pod 和名为“foo”的 service
        - kubectl delete pods,services -l name=myLabel                              # 删除具有 name=myLabel 标签的 pod 和 serivce
        - kubectl delete pods,services -l name=myLabel --include-uninitialized      # 删除具有 name=myLabel 标签的 pod 和 service，包括尚未初始化的
        - kubectl -n my-ns delete po,svc --all                                      # 删除 my-ns namespace 下的所有 pod 和 serivce，包括尚未初始化的
- 与运行中的 Pod 交互
    - kubectl (logs | exec | ...)
        - kubectl logs my-pod                                 # dump 输出 pod 的日志（stdout）
        - kubectl logs my-pod -c my-container                 # dump 输出 pod 中容器的日志（stdout，pod 中有多个容器的情况下使用）
        - kubectl logs -f my-pod                              # 流式输出 pod 的日志（stdout）
        - kubectl logs -f my-pod -c my-container              # 流式输出 pod 中容器的日志（stdout，pod 中有多个容器的情况下使用）
        - kubectl run -i --tty busybox --image=busybox -- sh  # 交互式 shell 的方式运行 pod
        - kubectl attach my-pod -i                            # 连接到运行中的容器
        - kubectl port-forward my-pod 5000:6000               # 转发 pod 中的 6000 端口到本地的 5000 端口
        - kubectl exec my-pod -- ls /                         # 在已存在的容器中执行命令（只有一个容器的情况下）
        - kubectl exec my-pod -c my-container -- ls /         # 在已存在的容器中执行命令（pod 中有多个容器的情况下）
        - kubectl top pod POD_NAME --containers               # 显示指定 pod 和容器的指标度量


## k8s物理结构
- 由多个物理主机（可以是虚拟机）构成集群，将多个底层物理主机的资源抽象出来组成一个平台进行编排管理
- 集群中所有的节点分成两类 : Master 和 Node，所有的容器中的应用程序都是运行在node中,而master是负责管理有多少个节点、每个节点上应该运行哪个或哪些容器的统一的控制中心。因此master被称为control plane（控制平面），node称为worker
- 多台master是为了冗余，避免但宕机，而不是用于负载均衡。而多个node就是为了负载均衡
- 支持多个容器引擎（docker...）,最好是docker，其他的容器引擎使用cri进行交互

## k8s master的组成
- 四个逻辑组件组成，每个组件都需要运行一个守护进程。
- API Server : 接受客户端请求的入口，检测语法规则，再存储在etcd中，controller watch API Server，有新任务时候由controller执行
- Scheduler : 调度器，监控集群所有节点的运算资源（cpu）和内存资源，根据资源情况分配给每个节点任务
- Controller : 控制器，执行用户命令，监控容器的存活情况，类似与大脑
- etcd (非k8s自己提供) : k-v存储系统，存储容器状态，保存所有集群数据，使用时需要为etcd数据提供备份计划

## k8s node的组成
- 每一个节点上有一个kublet进程，负责watch API Server，当分配给自己这个节点新的容器任务/容器状态变化时，就会调用Docker
- k8s虽然是容器编排系统，但是每个运行的容器，再k8s上都是进行重新封装过的，叫做pod，既k8s不会直接运行容器，而是运行pod，pod是容器的一层外壳。一个pod中可能存在多个容器，这多个容器当成一个原子来进行管理的。（既这些个容器都都只能运行在同一台节点之上），k8s运行的核心基本单元是pod，而不是容器。

## pod controller
- k8s运行的原子单元并非是容器，而是pod组件
- k8s最小的运行单元是pod组件，k8s对这些pod的创建、管理是由pod controller完成的，而非用户本身。我们告诉控制器我们需要的pod信息，其他的由控制器完成
- 由于pod的销毁/创建是动态的，在这样的动态场景中，我们需要依赖于服务注册、服务发现的方式来完成动态服务的衔接。k8s在pod和客户端之间添加了一个中间层service，来解决pod动态变动引起的ip的改变的问题，如果有多个pod，这个service还可以实现负载均衡的效果。可以理解service为pod的代理，当service创建完成之后也会有一个ip。 
- 最基础的pod控制器: Deployment

## k8s 对象的介绍

- 在Kubernetes中，几乎一切都是对象。常见的对象包括：Node，Pod，Deployment，ReplicationController， ReplicaSet等等。我们通常通过在描述文件中指定kind来创建不同种类的对象。Kubernetes通过etcd存储我们创建的对象，从而使应用按照你期望的方式稳定运行在容器中。
- Kubernetes对象本质上一种用于持久化的实体，Kubernetes使用这些持久化实体来描述一个集群。通常一个Kubernetes对象可以包含以下信息：
    - 需要运行的应用以及运行在哪些node上
    - 应用可以使用哪些资源
    - 应用运行时的一些配置，例如重启策略，升级以及容错性
- 每一个Kubernetes对象都包含两个属性：对象定义（Spec）和对象状态（Status）。Spec包含了你期望对象应有的状态，一般通过.yaml文件来描述。Status是该对象当前实际的状态。Kubernetes在任何状况下，它都会尽力确保对象的状态处于你所期望的状态。
- 举个例子，一个Deployment类型的对象，代表了你在集群上运行的一个应用。当你创建这个对象的时候，指定了replicas=3，这意味着你希望这个应用以拥有三个副本的状态下运行。如果其中一个副本由于程序问题崩溃了，那么当前应用的Status就是副本数为2，这与你期望的状态（Spec: replicas=3）不相同，因此Kubernetes会自动为你新建一个副本，从而使此应用的实际状态与你期望的状态相符。
- 创建一个Kubernetes对象通常分为两步：
    - 创建对象描述文件
    - 通过kubectl命令行接口创建对象
![avatar](/k8s/pic/deployment.png)


### k8s pod
- https://www.jianshu.com/p/ce71385e0370

### k8s replicaSet
- https://www.jianshu.com/p/fd8d8d51741e

### k8s label
- https://www.jianshu.com/p/cd6b4b4caaab
- Label以key/value键值对的形式附加到任何对象上，如Pod，Service，Node，RC（ReplicationController）/RS（ReplicaSet）等
- 在为对象定义好Label后，其他对象就可以通过Label来对对象进行引用, Label的最常见的用法便是通过spec.selector来引用对象
![avatar](/k8s/pic/label-selector.png)
- 我们通常使用metadata.labels字段，来为对象添加Label
- 带有Label的对象创建好之后，我们就可以通过Label Selector来引用这些对象。

### k8s service
- https://www.jianshu.com/p/c24fd0d132f6
- 通过ReplicaSet来创建一组Pod来提供具有高可用性的服务。虽然每个Pod都会分配一个单独的Pod IP，会有如下两个问题
    - Pod IP仅仅是集群内可见的虚拟IP，外部无法访问。
    - Pod IP会随着Pod的销毁而消失，当ReplicaSet对Pod进行动态伸缩时，Pod IP可能随时随地都会变化，这样对于我们访问这个服务带来了难度。
- Service可以看作是一组提供相同服务的Pod对外的访问接口。借助Service，应用可以方便地实现服务发现和负载均衡
- 它和其他Controller对象一样，通过Label Selector来确定一个Service将要使用哪些Pod
![avatar](/k8s/pic/service.png)
- 在Serive定义时，我们需要指定spec.type字段，这个字段拥有四个选项:
    - ClusterIP。默认值。给这个Service分配一个Cluster IP，它是Kubernetes系统自动分配的虚拟IP，因此只能在集群内部访问。
    - NodePort。将Service通过指定的Node上的端口暴露给外部。通过此方法，访问任意一个NodeIP:nodePort都将路由到ClusterIP，从而成功获得该服务。
    - LoadBalancer。在 NodePort 的基础上，借助 cloud provider 创建一个外部的负载均衡器，并将请求转发到 <NodeIP>:NodePort。此模式只能在云服务器（AWS等）上使用。
    - ExternalName。将服务通过 DNS CNAME 记录方式转发到指定的域名（通过 spec.externlName 设定）。需要 kube-dns 版本在 1.7 以上。

### k8s deployment
- https://www.jianshu.com/p/029661f38674



