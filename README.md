# p1-exporter

## Command line arguments

`-listen=<host>:<port>` Set the listen host/port. Default: `:9929`.
`-device=<device>` Specify the device associated to the RS232R device. Default: `/dev/ttyUSB0`.
`-baud=<baud rate>` Baud rate. Default: `115200`.
`-loglevel=<level>` Set the expected log level (error, warn, info, debug, trace). Default: `warn`.

const defaultListenAddress = ":9929"       // --listen=
const defaultSerialDevice = "/dev/ttyUSB0" // --device=
const defaultSerialBaud = 115200           // --baud=
const defaultLogLevel = "warn"             // --loglevel=

## OBIS Codes
List of OBIS code that are used by the Belgian (Fluvius) digital meters
```
OBIS CODE	BETEKENIS
0-0:96.1.4	ID
0-0:96.1.1	Serienummer van de elektriciteitsmeter (in ASCII hex)
0-0:1.0.0	Timestamp van de telegram
1-0:1.8.1	Tarief 1 (dag) – totaal verbruik
1-0:1.8.2	Tarief 2 (nacht) – totaal verbruik
1-0:2.8.1	Tarief 1 (dag) – totale injectie
1-0:2.8.2	Tarief 2 (nacht) – totale injectie
0-0:96.14.0	Huidig tarief (1=dag,2=nacht)
1-0:1.7.0	Huidig verbuik op alle fases
1-0:2.7.0	Huidige injectie op alle fases
1-0:21.7.0	L1 huidig verbruik
1-0:41.7.0	L2 huidig verbruik
1-0:61.7.0	L3 huidig verbruik
1-0:22.7.0	L1 huidige injectie
1-0:42.7.0	L2 huidige injectie
1-0:62.7.0	L3 huidige injectie
1-0:32.7.0	L1 spanning
1-0:52.7.0	L2spanning
1-0:72.7.0	L3spanning
1-0:31.7.0	L1 stroom
1-0:51.7.0	L2 stroom
1-0:71.7.0	L3 stroom
0-0:96.3.10	Positie schakelaar elektriciteit
0-0:17.0.0	Max. toegelaten vermogen/fase
1-0:31.4.0	Max. toegelaten stroom/fase
0-0:96.13.0	Bericht
0-1:24.1.0	Andere toestellen op bus
0-1:96.1.1	Serienummer van de aardgasmeter (in ASCII hex)
0-1:24.4.0	Positie schakelaar aardgas
0-1:24.2.3	Data van de aardgasmeter (timestamp) (waarde)
```