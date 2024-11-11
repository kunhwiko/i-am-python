### ECS
---
##### Elastic Container Service (ECS)
```
ECS is a managed container orchestrator service. ECS task definitions define what containers to deploy, IAM roles to assume,
resource requirements, logging settings, and volume mounts (e.g. EFS). Tasks can then be launched based off task definitions. 
```

##### ECS Service
```
a) ECS services help to run and maintain a specified number of instances of a task definition. If a task fails or stops,
   the service scheduler launches another instance of the task definition to replace it. Services are typically used to 
   deploy long-running tasks. Otherwise, short-term tasks can be deployed as a standalone without a service.
b) When deploying a task, VPC, load balancers, replica count, and security groups (i.e. each task has its own ENI and 
   can have exclusive security groups) can be configured.
c) Autoscaling can be enabled for services to increase the number of ECS tasks. Scaling can be done based on CPU/memory
   utilization of tasks, request count on load balancers, or on CloudWatch metrics. The ECS autoscaler can also be paired
   with ASGs to increase the number of underlying EC2 instances accordingly.
```

##### Launch Type
```
a) EC2 Launch Type: EC2 instances are provisioned and maintained by the user. Each instance must have the ECS agent 
   installed to be registered to the ECS cluster. EC2 instance profiles are used to grant permissions to the ECS agent
   (e.g. pull from registry, ship CloudWatch logs, reference Secrets Manager).
b) Fargate Launch Type: Fargate manages the provisioning of instances based on ECS task definitions and resource 
   requirements needed.
```

### EKS
---
##### Elastic Kubernetes Service (EKS)
```
EKS is AWS's managed Kubernetes service and leverages AWS infrastructure (e.g. EC2 instances, VPC networking, IAM). EKS
can be launched with EC2 instances or Fargate.
```

##### IAM Roles for Service Accounts (IRSA)
```
IRSA is an EKS native way to allow pods to assume IAM roles through service account annotations. It moves away from 
having to define IAM roles at the node level (i.e. EC2 instance).

Step 1) An OIDC provider for the cluster must be created with a public endpoint. This OIDC provider authenticates and 
        issues new JWT tokens that are projected into pods via service accounts.
Step 2) An IAM role must be created with a trust relationship to the OIDC provider. The trust policy indicates that any
        service authenticated by the OIDC provider can be trusted.
Step 3) Service accounts should be annotated with the ARN of the IAM role to assume. A mutating webhook called the Pod 
        Identity Webhook will inject environment variables representing the IAM Role ARN and path to the pod's service 
        account token.
Step 4) When a pod makes an AWS API call, the request is sent to STS to evaluate the trust policy of the IAM role. STS 
        verifies with the public OIDC endpoint to verify that the pod's JWT token is issued by the same OIDC provider.
Step 5) If verified, STS issues temporary credentials which effectively assigns the IAM permissions to the pod.
```

##### Karpenter
```
Karpenter is an open-source, flexible, high-performance Kubernetes cluster autoscaler. Compared to the traditional 
autoscaler, Karpenter provides the following:
  * No requirement to pre-specify AWS autoscaling groups or node pools. Karpenter has the ability to observe the
    aggregate resource requests of unscheduled pods by looking at Kubernetes events and will launch instances appropriate 
    for the workloads.
  * Faster node provisioning due to direct interaction with EC2 fleet APIs (i.e. no interaction with ASGs).
  * Provides a consolidation feature to move workloads from underutilized nodes to other nodes.
  * Provides a feature for node recycling to keep nodes up-to-date.
```

### App Runner
---
##### App Runner
```
App Runner is a service that simplifies deploying web applications. After users configure resource requirements, autoscaling,
health checks, and images, App Runner will deploy the application and manage the underlying infrastructure. This simplifies
availability, load balancing, encryption, and VPC access to other services.
```