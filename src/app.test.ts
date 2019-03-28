import got from "got";

const uri = "http://localhost:8080";

describe("check headers", () => {
  test("javascript", async () => {
    const { headers } = await got(`${uri}/sample.js`);

    expect(headers).toMatchObject({
      "cache-control": "max-age=31536000, public",
      "content-type": "application/javascript"
    });
  });

  test("service-worker.js", async () => {
    const { headers } = await got(`${uri}/service-worker.js`);

    expect(headers).toMatchObject({
      "cache-control": "no-cache",
      "content-type": "application/javascript"
    });
  });

  test("css", async () => {
    const { headers } = await got(`${uri}/sample.css`);

    expect(headers).toMatchObject({
      "cache-control": "max-age=31536000, public",
      "content-type": "text/css"
    });
  });

  test("asset-manifest.json", async () => {
    const { headers } = await got(`${uri}/asset-manifest.json`);

    expect(headers).toMatchObject({
      "cache-control": "no-cache",
      "content-type": "application/json"
    });
  });

  test("manifest.json", async () => {
    const { headers } = await got(`${uri}/manifest.json`);

    expect(headers).toMatchObject({
      "cache-control": "no-cache",
      "content-type": "application/json"
    });
  });

  test("index.html", async () => {
    const { headers } = await got(`${uri}`);

    expect(headers).toMatchObject({
      "cache-control": "no-cache",
      "content-type": "text/html"
    });
  });

  // Check we return idnex.html for routes
  test("index.html/some/path", async () => {
    const { headers } = await got(`${uri}/some/path`);

    expect(headers).toMatchObject({
      "cache-control": "no-cache",
      "content-type": "text/html"
    });
  });

  // Visual check for confirm this is not in the access logs
  test("/health", async () => {
    const { headers } = await got(`${uri}/health`, {
      headers: {
        "user-agent": `kube-probe/13.0.1`
      }
    });

    expect(headers).toMatchObject({
      // Returns index.html by default
      "cache-control": "no-cache",
      "content-type": "text/html"
    });
  });

  test("png", async () => {
    const { headers } = await got(`${uri}/sample.png`);

    expect(headers).toMatchObject({
      "content-type": "image/png"
    });
  });
});
