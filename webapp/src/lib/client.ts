import type { Land } from "./types";

const BACKEND_URL = "192.168.1.10:8080";

export const getLands = async (): Promise<Land[]> => {
    const response = await fetch(`${BACKEND_URL}/land-registry`);
    return await response.json();
}