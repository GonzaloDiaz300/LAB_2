#Crea dockerfile e inicia contenedor para servidor OMS
docker-OMS:
	docker build -f Dockerfile.oms . -t containerized_oms:latest
	docker run --rm --name oms-server -p 50051:50051 --network="host" containerized_oms:latest

#Genera la imagen de australia.
docker-continentes:
	docker build -f Dockerfile.australia . -t containerized_australia:latest
	docker run --rm --name australia-server -p 50055:50055 --network="host" containerized_australia:latest