import { Service } from "./service";
import type { APIResponse, Ping } from "@/models/APIResponse";

class PingService extends Service {
  ping(): Promise<APIResponse<Ping>> {
    return this.client.get("").json();
  }
}

export const pingService = new PingService("/api/ping")
