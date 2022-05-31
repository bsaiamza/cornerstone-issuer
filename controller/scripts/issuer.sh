#!/bin/bash
docker stop cornerstone-issuer-agent
docker rm cornerstone-issuer-agent

docker run -p 10000:10000 -p 10001:10001  --name cornerstone-issuer-agent --network=bridge -it bcgovimages/aries-cloudagent:py36-1.16-1_0.7.3 start \
-l 'Iamza Cornerstone Issuer' \
-it http 0.0.0.0 10000 \
-ot http \
--admin 0.0.0.0 10001 \
--admin-insecure-mode \
-e http://172.17.0.2:10000/ \
--genesis-url http://172.20.0.1:9000/genesis \
--seed IamzaCornerstoneIssuer0000000000  \
--wallet-type indy \
--wallet-name issuer_wallet \
--wallet-key issuer_secret \
--log-level 'info'  \
--auto-provision \
--auto-accept-invites \
--auto-accept-requests \
--auto-ping-connection 