-- docker exec -it infra_cassandra_1 cqlsh


CREATE KEYSPACE IF NOT EXISTS store WITH REPLICATION = { 
    'class' : 'SimpleStrategy',
    'replication_factor' : '1' 
};

CREATE TABLE IF NOT EXISTS store.message (
    create_ts timestamp PRIMARY KEY,
    message text
);