#!/usr/bin/env sh

CONSUL=${CONSUL_URL}
MGOURL=${MONGO_URL}
KAFKAURL=${KAFKA_URL}

L_CONSUL=${CONSUL//\//\\\/}
L_MGOURL=${MGOURL//\//\\\/}
L_KAFKAURL=${KAFKAURL//\//\\\/}

echo "consul url: ${L_CONSUL}"
echo "mongourl url: ${L_MGOURL}"
echo "kafka url: ${L_KAFKAURL}"

echo "set mongo and consul url"
find /home/game/bin/ -type f -name "*.json" -exec sed -i "s/MONGO_URL/${L_MGOURL}/g" {} \;
find /home/game/bin/ -type f -name "*.json" -exec sed -i "s/CONSUL_URL/${L_CONSUL}/g" {} \;
find /home/game/bin/ -type f -name "*.json" -exec sed -i "s/KAFKA_URL/${L_KAFKAURL}/g" {} \;

chmod u+x /home/game/bin/gwapi
echo "start game gateway"
/home/game/bin/gwapi
