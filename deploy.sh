# requires local registry to run, use:  docker run -d -p 5000:5000 --name registry registry:2
# requires helm to be installed
# requires helm chart to be installed, use: helm install openesl-chart openeslchart/ --values openeslchart/values.yaml
docker build -t openesl .
docker image tag openesl localhost:5000/openesl
docker push localhost:5000/openesl
helm upgrade openesl-chart openeslchart/ --values openeslchart/values.yaml