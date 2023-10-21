#Crea dockerfile e inicia contenedores para servidores Datanodes
docker-datanode:
	docker build -f Dockerfile.datanode2 . -t containerized_datanode2:latest
	docker run --rm --name datanode2-server -p 50054:50054 --network="host" containerized_datanode2:latest
#Genera todas las imagenes de los servidores continentes.
docker-continentes:
	docker build -f Dockerfile.latinoamerica . -t containerized_latinoamerica:latest
	docker run --rm --name latinoamerica-server -p 50058:50058 --network="host" containerized_latinoamerica:latest