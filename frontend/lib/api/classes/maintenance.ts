import { BaseAPI, route } from "../base";
import type {
  MaintenanceEntry,
  MaintenanceEntryWithDetails,
  MaintenanceEntryUpdate,
  MaintenanceFilterStatus,
} from "../types/data-contracts";
import type { AttachmentTypes } from "../types/non-generated";
import type { Requests } from "~~/lib/requests";
import { AttachmentsAPI } from "./attachments";

export interface MaintenanceFilters {
  status?: MaintenanceFilterStatus;
}

export class MaintenanceAPI extends BaseAPI {
  attachments: AttachmentsAPI;

  constructor(http: Requests, token: string) {
    super(http, token);
    this.attachments = new AttachmentsAPI(http);
  }

  getAttachments(id: string) {
    return this.attachments.getAttachments("maintenance", id);
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

  uploadAttachment(id: string, file: File, filename: string, type?: AttachmentTypes, primary?: boolean) {
    return this.attachments.addAttachment("maintenance", id, file, filename, type, primary);
  }

  deleteAttachment(attachmentId: string) {
    return this.attachments.deleteAttachment(attachmentId);
  }

  updateAttachment(
    attachmentId: string,
    payload: { title?: string; type?: AttachmentTypes; primary?: boolean }
  ) {
    return this.attachments.updateAttachment(attachmentId, payload);
  }

}
