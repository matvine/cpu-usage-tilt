# -*- mode: Python -*-

docker_build('cpu-cores-image', './app', dockerfile='deployments/Dockerfile')

k8s_yaml(
  ['deployments/deployment.yaml', 'deployments/deployment-maxprocs.yaml'], 
  allow_duplicates=True
)

k8s_resource('cpu-cores-no-max-procs',
  port_forwards='8000:8009'
)


k8s_resource('cpu-cores-max-procs',
  port_forwards='8001:8009'
)

load('ext://helm_remote', 'helm_remote')

helm_remote('grafana', 
  repo_url='https://grafana.github.io/helm-charts',
  set=["service.port=8080"]
)

k8s_resource(workload='grafana', port_forwards='8080:8080')

k8s_yaml(['cadvisor/daemonset.yaml', 'cadvisor/namespace.yaml', 'cadvisor/serviceaccount.yaml'])
k8s_resource('cadvisor', 
  port_forwards='8008:8080'
)