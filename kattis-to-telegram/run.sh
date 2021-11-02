# Allow 'default' user in huginn container to rwx
mkdir -pv -m 777 db
docker-compose up -d
