#Crea dockerfile e inicia contenedor para servidor OMS
docker-OMS:
	docker build -f Dockerfile.oms . -t containerized_oms:latest
	docker run --rm --name oms-server -p 50051:50051 --network="host" containerized_oms:latest

#Crea dockerfile e inicia contenedor para servidor ONU
docker-ONU:
	docker build -f Dockerfile.onu . -t containerized_onu:latest
	docker run --rm --name onu-server -p 50052:50052 --network="host" -i containerized_onu:latest

#Crea dockerfile e inicia contenedores para servidores Datanodes
docker-datanode:
	docker build -f Dockerfile.datanode1 . -t containerized_datanode1:latest
	docker build -f Dockerfile.datanode2 . -t containerized_datanode2:latest

	docker run --rm --name datanode1-server -p 50053:50053 --network="host" containerized_datanode1:latest
	docker run --rm --name datanode2-server -p 50054:50054 --network="host" containerized_datanode2:latest

docker-datanode1:
	docker build -f Dockerfile.datanode1 . -t containerized_datanode1:latest
	docker run --rm --name datanode1-server -p 50053:50053 --network="host" containerized_datanode1:latest

docker-datanode2:
	docker build -f Dockerfile.datanode2 . -t containerized_datanode2:latest
	docker run --rm --name datanode2-server -p 50054:50054 --network="host" containerized_datanode2:latest

#Genera todas las imagenes de los servidores continentes.
docker-continentes:
	docker build -f Dockerfile.australia . -t containerized_australia:latest
	docker build -f Dockerfile.asia . -t containerized_asia:latest
	docker build -f Dockerfile.europa . -t containerized_europa:latest
	docker build -f Dockerfile.latinoamerica . -t containerized_latinoamerica:latest

	docker run --rm --name australia-server -p 50055:50055 --network="host" containerized_australia:latest
	docker run --rm --name asia-server -p 50056:50056 --network="host" containerized_asia:latest
	docker run --rm --name europa-server -p 50057:50057 --network="host" containerized_europa:latest
	docker run --rm --name latinoamerica-server -p 50058:50058 --network="host" containerized_latinoamerica:latest

docker-continentes-australia:
	docker build -f Dockerfile.australia . -t containerized_australia:latest
	docker run --rm --name australia-server -p 50055:50055 --network="host" containerized_australia:latest

docker-continentes-asia:
	docker build -f Dockerfile.asia . -t containerized_asia:latest
	docker run --rm --name asia-server -p 50056:50056 --network="host" containerized_asia:latest

docker-continentes-europa:
	docker build -f Dockerfile.europa . -t containerized_europa:latest
	docker run --rm --name europa-server -p 50057:50057 --network="host" containerized_europa:latest

docker-continentes-latinoamerica:
	docker build -f Dockerfile.latinoamerica . -t containerized_latinoamerica:latest
	docker run --rm --name latinoamerica-server -p 50058:50058 --network="host" containerized_latinoamerica:latest