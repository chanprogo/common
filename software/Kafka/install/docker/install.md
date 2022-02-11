
`docker pull wurstmeister/zookeeper:latest`  
`docker pull wurstmeister/kafka:2.12-2.3.1`  



```
docker run -d --name zookeeper \
-p 2181:2181 \
-v /home/chan/Desktop/kafka-test/zookeeper/data:/opt/zookeeper-3.4.13/data \
-v /etc/localtime:/etc/localtime:ro \
-t wurstmeister/zookeeper
```  


```
docker run -d --name kafka \
-p 9092:9092 \
-v /home/chan/Desktop/kafka-test/kafka/logs:/opt/kafka/logs \
-v /home/chan/Desktop/kafka-test/kafka/kafka-logs:/kafka/kafka-logs \
-v /etc/localtime:/etc/localtime:ro \
-e KAFKA_BROKER_ID=0 \
-e KAFKA_ZOOKEEPER_CONNECT=172.16.201.231:2181 \
-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://47.113.189.98:9092 \
-e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
-e KAFKA_LOG_DIRS=/kafka/kafka-logs \
-t wurstmeister/kafka:2.12-2.3.1`  
```

`vi opt/kafka/config/server.properties`  
