### dockerの中に入る
docker container exec -it 42tokyo-road-to-dojo-go_mysql_1 bash

### mysqlにログイン
mysql -uroot -pca-tech-dojo

### Re-Initialize
docker-compose down
docker volume rm 42tokyo-road-to-dojo-go_db-data
docker-compose up -d mysql
