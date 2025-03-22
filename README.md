## sx-evo-debug

A lightweight Dockerized toolkit for debugging CAN traffic on the DOBISS SX Evolution

```bash
Starting CAN interface setup...
Setting up CAN interface...
Bringing up can0 interface...
Setting txqueuelen to 1000...
CAN interface setup completed successfully
Starting sx-evo-debug...
2025/03/22 11:44:04 INFO received can frame module=G output=5 id=16842496 data=00470400
2025/03/22 11:44:04 INFO received can frame module=E output=7 id=16842496 data=00450600
2025/03/22 11:44:05 INFO received can frame module=G output=1 id=16842496 data=00470000
2025/03/22 11:44:05 INFO received can frame module=E output=7 id=16842496 data=00450600
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
