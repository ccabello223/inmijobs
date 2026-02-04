import ky from "ky";
import type { KyInstance } from "ky";

export class Service {
  protected basePath: string;
  protected client: KyInstance;

  constructor(basePath: string) {
    try {
      const url = new URL("http://localhost:3000" + basePath);
      this.basePath = url.toString();
      this.client = ky.extend({
        prefixUrl: url.toString(),
        credentials: "include",
      })
      console.log("Service created successfully:", this.basePath)
    } catch (error) {
      console.error("Error creating a new service: ", error)
      throw error;
    }
  }
}
