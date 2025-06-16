import { BaseAPI, route } from "../base";
import type {
  MaintenanceEntry,
  MaintenanceEntryWithDetails,
  MaintenanceEntryUpdate,
  MaintenanceFilterStatus,
} from "../types/data-contracts";
import type { MaintenanceEntryAttachment } from "../types/data-contracts";
import type { AttachmentTypes } from "../types/non-generated";
import type { Requests } from "~~/lib/requests";

export interface MaintenanceFilters {
  status?: MaintenanceFilterStatus;
}

export class MaintenanceAPI extends BaseAPI {
  attachments: MaintenanceAttachmentsAPI;

  constructor(http: Requests, token: string) {
    super(http, token);
    this.attachments = new MaintenanceAttachmentsAPI(http);
  }
  getAttachments(id: string) {
    return this.http.get<{ data: MaintenanceEntryAttachment[] }>({
      url: route(`/maintenance/${id}/attachments`),
    });
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

  uploadAttachments(id: string, form: FormData) {
    return this.http.post<FormData, MaintenanceEntryAttachment>({
      url: route(`/maintenance/${id}/attachments`),
      body: form,
      headers: { "Content-Type": "multipart/form-data" },
    });
  }

  /**
   * Verwijder een bijlage van een maintenance entry
   */
  deleteAttachment(id: string, attachmentId: string) {
    return this.http.delete<{ attachments: MaintenanceEntryAttachment[] }>({
      url: route(`/maintenance/${id}/attachments/${attachmentId}`),
    });
  }

  /**
   * Update een bijlage van een maintenance entry
   */
  updateAttachment(
    id: string,
    attachmentId: string,
    payload: { title?: string; type?: string; primary?: boolean }
  ) {
    return this.http.put<
      typeof payload,
      { attachments: MaintenanceEntryAttachment[] }
    >({
      url: route(`/maintenance/${id}/attachments/${attachmentId}`),
      body: payload,
    });
  }
}

export class MaintenanceAttachmentsAPI extends BaseAPI {
  addAttachment(
    id: string,
    file: File | Blob,
    filename: string,
    type: AttachmentTypes | null = null,
    primary?: boolean
  ) {
    const formData = new FormData();
    formData.append("file", file);
    if (type) {
      formData.append("type", type);
    }
    formData.append("name", filename);
    if (primary !== undefined) {
      formData.append("primary", primary.toString());
    }

    return this.http.post<FormData, { attachments: MaintenanceEntryAttachment[] }>({
      url: route(`/maintenance/${id}/attachments`),
      body: formData,
      headers: { "Content-Type": "multipart/form-data" },
    });
  }

}
