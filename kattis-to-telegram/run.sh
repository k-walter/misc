# Allow 'default' user in huginn container to rwx
mkdir -pv -m 777 mysql-data
docker-compose up -d
