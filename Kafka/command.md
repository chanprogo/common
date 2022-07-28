
单机方式：创建一个主题   
`bin/kafka-topics.sh --create --zookeeper 172.16.201.231:2181 --replication-factor 1 --partitions 1 --topic mykafka`  

运行一个生产者  
`bin/kafka-console-producer.sh --broker-list localhost:9092 --topic mykafka`  

运行一个消费者  
`bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic mykafka --from-beginning`   
