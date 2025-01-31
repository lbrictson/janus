# Installation

## Docker Install

Running Janus with Docker is the easiest way to get up and running.

```bash
cd ~/some-directory-where-you-want-to-store-data/
docker run -it -v $PWD:/data/ -e JANUS_URL=https://yourdomain.com -p 8080:8080 -p 8081:8081 lbrictson/janus:latest-slim
# Access at http://localhost:8080/
# Username: admin@localhost
# Password: ChangeMeBeforeUse1234!
# Replace https://yourdomain.com with the domain your plan to host janus at
```

Note that when using docker the sqlite database is stored in `/data/janus.db`.  To persist this data between
restart you need to either save it to a mappe folder like the example above or to a volume.

## Docker Compose Install

```dockerfile
services:
  janus:
    platform: linux/x86_64
    volumes:
      - janusdb:/data
    ports:
      - "8080:8080"
      - "8081:8081"
    image: lbrictson/janus:latest-slim
    environment:
      # JANUS_URL should be the URL where you plan to host Janus
      - JANUS_URL=https://yourdomain.com
volumes:
  janusdb:
```

## Binary Install

1. Download the latest release at: https://github.com/lbrictson/janus/releases
2. Decompress the executable: `tar -xvf janus_Linux_arm64.tar.gz`
3. Run Janus, make sure to set the URL to where you plan to access it.

```bash
 JANUS_URL=https://mydomain.com ./janus
```

Once you have Janus running you can access the web UI at http://localhost:8080/

The default username and password are
```bash
Username: admin@localhost
Password: ChangeMeBeforeUse1234!
```

You should change the password after logging in.

