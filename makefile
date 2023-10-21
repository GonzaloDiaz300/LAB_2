#Crea dockerfile e inicia contenedores para servidores Datanodes
docker-datanode:
	docker build -f Dockerfile.datanode1 . -t containerized_datanode1:latest
	docker run --rm --name datanode1-server -p 50053:50053 --network="host" containerized_datanode1:latest

#Genera la imagen de europa
docker-continentes:
	docker build -f Dockerfile.europa . -t containerized_europa:latest
	docker run --rm --name europa-server -p 50057:50057 --network="host" containerized_europa:latest