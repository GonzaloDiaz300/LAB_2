#Crea dockerfile e inicia contenedor para servidor ONU
docker-ONU:
	docker build -f Dockerfile.onu . -t containerized_onu:latest
	docker run --rm --name onu-server -p 50052:50052 --network="host" -i containerized_onu:latest
#Genera imagen de asia
docker-continentes:
	docker build -f Dockerfile.asia . -t containerized_asia:latest
	docker run --rm --name asia-server -p 50056:50056 --network="host" containerized_asia:latest