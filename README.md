## sx-evo-debug

A lightweight Dockerized toolkit for debugging CAN traffic on the DOBISS SX Evolution

```bash
Starting CAN interface setup...
Setting up CAN interface...
Bringing up can0 interface...
Setting txqueuelen to 1000...
CAN interface setup completed successfully
Starting sx-evo-debug...
2025/03/14 19:13:55 | Module A | Output 4 | ID 16842496 | 00 41 03 00 
2025/03/14 19:13:55 | Module A | Output 4 | ID 16842496 | 00 41 03 01
```

### Features

- [x] Supports your favorite CANable (or CANable clone)
- [x] Filter traffic on a specific module
- [x] Filter traffic on a specific output
- [ ] Only show traffic from the DOBISS webserver

### Build

```bash
docker build --tag sx-evo-debug .
```

### Usage

#### Log all CAN trafic from /dev/ttyACM1 (mounted as can0)

```bash
docker run --name sx-evo-debug --rm --privileged --volume /dev:/dev --net=host -e CAN_DEVICE=/dev/ttyACM1 -e CAN_INTERFACE=can0 sx-evo-debug
```

#### Only log CAN traffic from module E

```bash
docker run --name sx-evo-debug --rm --privileged --volume /dev:/dev --net=host -e CAN_DEVICE=/dev/ttyACM1 -e CAN_INTERFACE=can0 -e DOBISS_MODULE=E sx-evo-debug
```

#### Only log CAN traffic from module A, output 1

```bash
docker run --name sx-evo-debug --rm --privileged --volume /dev:/dev --net=host -e CAN_DEVICE=/dev/ttyACM1 -e CAN_INTERFACE=can0 -e DOBISS_MODULE=E -e DOBISS_OUTPUT=1 sx-evo-debug
```
