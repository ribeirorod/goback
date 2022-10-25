
# Directory to persist postgres data
mkdir -p ~/postgresql/mount

cd ./docker

# Build docker image 

docker build --no-cache -t img-backend-db ./

# Run postgres container
docker run -d \
	-e POSTGRES_PASSWORD=docker \
	--name=db \
	-p 9090:5432 \
	-v ~/private/repos/model/backend-go/data:/var/lib/postgresql/data \
	db:new