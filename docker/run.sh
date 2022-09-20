
# Directory to persist postgres data
mkdir -p ~/postgresql/mount

cd ./docker

# Build docker image 

docker build --no-cache -t img-backend-db ./

# Run postgres container
docker run -d \
	-e POSTGRES_PASSWORD=docker \
	--name backend-db \
	-p 9090:5432 \
	-v ~/postgresql/mount:/var/lib/postgresql/data \
	img-backend-db
