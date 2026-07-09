# AndroidLibXrayLite

## Build requirements
* JDK
* Android SDK
* Go
* gomobile

## Build instructions
1. `git clone [repo] && cd AndroidLibXrayLite`
2. `gomobile init`
3. `go mod tidy -v`
4. `gomobile bind -v -androidapi 24 -trimpath -ldflags='-s -w -buildid= -checklinkname=0' ./`

## Xray-core version pinning

This fork keeps the current AndroidLibXrayLite wrapper API used by v2rayNG, but does not force the upstream `autorepobot/xray-core` replacement by default. The default `go.mod` dependency is:

```text
github.com/xtls/xray-core v1.260327.0
```

Use the `Build` workflow manually to publish a custom AAR:

```text
release_tag = v26.3.27-v2rayng-r2
xray_core_version =
xray_core_replace =
```

Leaving `xray_core_version` and `xray_core_replace` empty uses the version pinned in `go.mod`. To test another core module version, set `xray_core_version`. To intentionally use a replacement module, set `xray_core_replace` in `module@version` form.
