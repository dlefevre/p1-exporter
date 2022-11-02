package config

// Overwritable defaults
const defaultListenAddress = ":9929"       // --listen=
const defaultSerialDevice = "/dev/ttyUSB0" // --device=
const defaultSerialBaud = 115200           // --baud=
const defaultLogLevel = "warn"             // --loglevel=

// These defaults are hardcoded and cannot be overwritten with flags
const defaultSerialSize = 8
const defaultSerialStopBits = 1
const defaultBackOffInterval = 10
