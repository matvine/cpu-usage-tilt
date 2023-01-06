# -*- mode: Python -*-

docker_build('maxprocs', './maxprocs', dockerfile='maxprocs/Dockerfile')
docker_build('nomaxprocs', './nomaxprocs', dockerfile='nomaxprocs/Dockerfile')


k8s_yaml(
  ['deployments/nomaxprocs.yaml', 'deployments/nomaxprocs_limit.yaml','deployments/maxprocs.yaml'], 
  allow_duplicates=True
)

k8s_resource('cpu-cores-no-max-procs',
  port_forwards='8000:8009'
)

k8s_resource('cpu-cores-no-max-procs-limit',
  port_forwards='8001:8009'
)

k8s_resource('cpu-cores-max-procs',
  port_forwards='8002:8009'
)


k8s_yaml(['cadvisor/daemonset.yaml', 'cadvisor/namespace.yaml', 'cadvisor/serviceaccount.yaml'])
k8s_resource('cadvisor', 
  port_forwards='8008:8080'
)

k8s_yaml(['prometheus/deployment.yaml','prometheus/namespace.yaml','prometheus/config-map.yaml','prometheus/clusterrole.yaml'])
k8s_resource('prometheus-deployment', port_forwards='9090:9090')