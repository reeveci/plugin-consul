# Reeve CI / CD - Consul Plugin

This is a [Reeve](https://github.com/reeveci/reeve) plugin for providing pipeline environment variables from a Consul KV store.

## Configuration

### Consul

An API token is required for this plugin.
It is recommended to use a token configured with minimal required access.

### Settings

Settings can be provided to the plugin through environment variables set to the reeve server.

Settings for this plugin should be prefixed by `REEVE_PLUGIN_CONSUL_`.

Settings may also be shared between plugins by prefixing them with `REEVE_SHARED_` instead.

- `ENABLED` - `true` enables this plugin
- `URL` (required) - Consul URL
- `TOKEN` (required) - Consul API Token
- `KEY_PREFIX` - Key prefix
- `PRIORITY` (default `1`) - Priority of all variables returned by this plugin
- `SECRET` - `true` marks all variables returned by this plugin as secret
