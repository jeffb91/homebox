import { BaseAPI, route } from "../base";
import type {
  MaintenanceEntry,
  MaintenanceEntryWithDetails,
  MaintenanceEntryUpdate,
  MaintenanceFilterStatus,
} from "../types/data-contracts";

export interface MaintenanceFilters {
  status?: MaintenanceFilterStatus;
}

export class MaintenanceAPI extends BaseAPI {
  getAttachments(id: string): { data: any; error: any; } | PromiseLike<{ data: any; error: any; }> {
    throw new Error("Method not implemented.");
  }
  getAll(filters: MaintenanceFilters) {
    return this.http.get<MaintenanceEntryWithDetails[]>({
      url: route(`/maintenance`, { status: filters.status?.toString() }),
    });
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/maintenance/${id}`) });
  }

  update(id: string, data: MaintenanceEntryUpdate) {
    return this.http.put<MaintenanceEntryUpdate, MaintenanceEntry>({
      url: route(`/maintenance/${id}`),
      body: data,
    });
  }
}
