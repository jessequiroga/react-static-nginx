# Read Me

This simple contains a common setup used with apps deployed using create-react-app.
It also caters for service worker scenarios

## Updating

If you need to update this image make sure you bump the `REVISION` variable in the `.gitlab.ci.yml` file
This is because kubernetes does not re-download images so we need a new tag

If you update the nginx version also udpate the `NGINX_VERSION` variable

## Testing

    docker-compose up

Confirm that the sample files has the correct headers and that redirecting works as expected

- index.html should not be cached
- css should be cached for one year
- js should be cached for one year
- `service-worker.js` should never be cached
- /somepath should return index.html
- TODO add more

## TODO

Automate header testing

## Implementation

This image expects the following tags to be present in the head of the index.html page.

    <meta property="app:server-uri" content="__APP_SERVER_URI__">
    <meta property="app:stage" content="__APP_STAGE__">

These properties are expected to exist and are available from the `@stackworx/react/ConfigService`
