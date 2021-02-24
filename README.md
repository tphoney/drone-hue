# drone-hue

A plugin to get notifications via your light bulbs.

## Usage

There are a few settings used by the plugin that changes its behavior.

* hub_ip **(required)** IP address of the philips hue hub.
* hub_token **(required)** IP address of the philips hue hub
* target_type (optional) whether or not you want to target a bulb or a light or bulbs. Valid values are (lights, groups) defaults to groups.
* target (optional) the id of the target. defaults to 0, 0 targets everything in light/group.
* payload (optional) what you want the light/group to do when there are **no failed steps**. defaults to `{"alert":"lselect"}` which sets the bulbs to flash for 15 seconds.
* failed_payload (optional) what you want the light/group to do when there are **failed steps**. defaults to `{"alert":"lselect"}` which sets the bulbs to flash for 15 seconds.

The following .drone.yml builds some golang and runs the drone_hue step. If there is any failure in the golang build step it will change hue group `5` to the colour red and flash for 15 seconds.

```yaml
kind: pipeline
name: default

steps:
- name: test
  image: golang
  commands:
  - go test ./...
  - go build ./...
- name: drone_hue
  image: tphoney/drone-hue
  pull: if-not-exists
  settings:
    log_level: debug
    hub_ip: 192.168.0.115
    hub_token: zpSMgrRfIA-JC8BQQrqGquobI-SsT0v7hsm5gV7R
    target: 5
    fail_payload: '{"alert":"lselect", "hue":0}'
  when:
    status:
    - success
    - failure
```

NB The `when` clause means this step will run even if other steps fail, otherwise we would not run the `drone_hue` step.

For more information on how to setup your philips hub look [here](https://developers.meethue.com/develop/get-started-2/) and for further information about the philips api [here](https://developers.meethue.com/develop/hue-api/groupds-api/)

## Building

Build the plugin binary:

```text
scripts/build.sh
```

Build the plugin image:

```text
docker build -t tphoney/drone-hue -f docker/Dockerfile .
```

## Testing

Execute the plugin from your current working directory:

```text
docker run --rm -e PLUGIN_HUB_IP=192.168.0.115 -e PLUGIN_HUB_TOKEN=zpSMgrRfIA-JC8BQQrqGquobI-SsT0v7hsm5gV7R \
  -e DRONE_COMMIT_SHA=8f51ad7884c5eb69c11d260a31da7a745e6b78e2 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_BUILD_NUMBER=43 \
  -e DRONE_BUILD_STATUS=success \
  -w /drone/src \
  -v $(pwd):/drone/src \
  tphoney/drone-hue
```
