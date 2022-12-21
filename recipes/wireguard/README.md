# Wireguard

#### Installation

```bash
mkdir wireguard
```

```bash
wget https://raw.githubusercontent.com/ZanMax/homelab/main/wireguard/docker-compose.yml
```
> Change option in docker-compose
- TZ - timezone
- SERVERURL - set your public IP
- PEERS - number of clients
- /path/to/appdata/config - path to your folver with wireguard ( for example /home/dev/wireguard/config )


```bash
docker-compose up -d
```

#### Get client config

show config file

```bash
docker exec -it wireguard cat /config/peer1/peer1.conf
```
or
```bash
cat /home/dev/wireguard/config/peer1/peer1.conf
```

show PNG file as QR code
```bash
docker exec -it wireguard /app/show-peer 1
```
or you can find generated QR
/home/dev/wireguard/config/peer1/peer1.png

#### Generate QR from config

```bash
qrencode -r peer1.conf -t PNG -o daha_vpn.png
```