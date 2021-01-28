# vector bluetooth

Primarily, this is meant to be a set of libraries to connect to vector via BLE.  It just so happens that there's also a CLI that utilizes them built in.

## Prerequisites.

1.  You'll need to install [libsodium](https://libsodium.gitbook.io/doc/) for your OS.  Specific instructions vary by OS

## Running on Linux

To build the binary, run
```
$ make build
```

You'll have to do something like this on the built binary

```sh
$ sudo setcap 'cap_net_raw,cap_net_admin+eip' /path/to/file
```

## Running on OS X

Completely untested, but let me know!

## Running on Windows

Doesn't work yet...

### CLI

There are more features in the library than are currently available in the CLI, but here's whats currently available

|  Name | Description  |
| ------------ | ------------ |
|  scan | runs a BLE scan and displays an appropriate list of devices  |
|  connect | connect to a vector via ID (displayed in the scan)  |
|  authorize | performs a cloud authorization (but you'll need to find your token!)  |
|  configure | allows you to make/change configuration  |
|  get-status | displays the status of your vector  |
|  wifi-scan | scan for a list of available wifi networks  |
|  wifi-connect | connect to a wifi network  |
|  wifi-forget | tell vector to forget a wifi network |
|  wifi-ip | display IP information of  the robot |
|  ota-start | perform an OTA code download  |
|  ota-cancel | stops an in-progress OTA download |

### Project status

This is an early release, and not all features are implemented in either the CLI or in the library itself.

#### TODO:
1.  wifi-ip (get vectors IP) [issue](https://github.com/digital-dream-labs/vector-bluetooth/issues/1)
2.  wifi-ap (set up vector in AP mode) [issue](https://github.com/digital-dream-labs/vector-bluetooth/issues/2)
3.  wifi-forget (forget one or all wifi networks) [issue](https://github.com/digital-dream-labs/vector-bluetooth/issues/3)
4.  logs (download lotgs from vector) [issue](https://github.com/digital-dream-labs/vector-bluetooth/issues/4)

