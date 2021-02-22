# drone-hue

A plugin to get notifications via your light bulbs.

## Usage

There are a few parameters used by the plugin that changes its behavior.

* PLUGIN_HUB_IP **(required)** IP address of the philips hue hub
* PLUGIN_HUB_TOKEN **(required)** IP address of the philips hue hub
* PLUGIN_TARGET_TYPE (optional) whether or not you want to target a bulb or a light or bulbs. Valid values are (lights, groups) defaults to groups
* PLUGIN_TARGET (optional) the id of the target. defaults to 0
* PLUGIN_PAYLOAD (optional) what you want the bulb to do. defaults to `{"alert":"lselect"}` which sets the bulbs to flash for 15 seconds

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
docker run --rm \
  -e DRONE_COMMIT_SHA=8f51ad7884c5eb69c11d260a31da7a745e6b78e2 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_BUILD_NUMBER=43 \
  -e DRONE_BUILD_STATUS=success \
  -w /drone/src \
  -v $(pwd):/drone/src \
  tphoney/drone-hue
```
