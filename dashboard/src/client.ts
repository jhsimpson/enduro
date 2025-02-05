import * as api from "./openapi-generator";
import { usePackageStore, PackageMonitorResponseBody } from "./stores/package";

export interface Client {
  package: api.PackageApi;
  storage: api.StorageApi;
  connectPackageMonitor: () => void;
}

function getPath(): string {
  const base = "/api";
  const location = window.location;
  const path =
    location.protocol +
    "//" +
    location.hostname +
    (location.port ? ":" + location.port : "") +
    base +
    (location.search ? location.search : "");

  return path.replace(/\/$/, "");
}

function storageServiceDownloadURL(aipId: string): string {
  return (
    getPath() +
    `/storage/{aip_id}/download`.replace(
      `{${"aip_id"}}`,
      encodeURIComponent(aipId)
    )
  );
}

function getWebSocketURL(): string {
  let url = getPath();

  if (url.startsWith("https")) {
    url = "wss" + url.slice("https".length);
  } else if (url.startsWith("http")) {
    url = "ws" + url.slice("http".length);
  }

  return url;
}

function connectPackageMonitor() {
  const store = usePackageStore();
  const url = getWebSocketURL() + "/package/monitor";
  const socket = new WebSocket(url);
  socket.onmessage = (event: MessageEvent) => {
    const body = JSON.parse(event.data) as PackageMonitorResponseBody;
    store.handleEvent(body);
  };
  socket.onclose = (event: CloseEvent) => {
    // tslint:disable-next-line:no-console
    console.log("Enduro WebSocket client closed", event.code);
  };

  // tslint:disable-next-line:no-console
  console.log("Enduro WebSocket client created", url);
}

function createClient(): Client {
  const path = getPath();
  const config: api.Configuration = new api.Configuration({ basePath: path });

  // tslint:disable-next-line:no-console
  console.log("Enduro client created", path);

  return {
    package: new api.PackageApi(config),
    storage: new api.StorageApi(config),
    connectPackageMonitor,
  };
}

const client = createClient();

export { api, client, storageServiceDownloadURL };
