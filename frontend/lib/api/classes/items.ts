import { BaseAPI, route } from "../base";
import { parseDate } from "../base/base-api";
import type {
  ItemCreate,
  ItemOut,
  ItemPatch,
  ItemPath,
  ItemSummary,
  ItemUpdate,
  MaintenanceEntry,
  MaintenanceEntryCreate,
  MaintenanceEntryWithDetails,
} from "../types/data-contracts-old";
import type { AttachmentTypes, ItemSummaryPaginationResult } from "../types/non-generated";
import type { MaintenanceFilters } from "./maintenance";
import type { Requests } from "~~/lib/requests";
import { AttachmentsAPI as GenericAttachmentsAPI } from "./attachments";

export type ItemsQuery = {
  orderBy?: string;
  includeArchived?: boolean;
  page?: number;
  pageSize?: number;
  locations?: string[];
  labels?: string[];
  negateLabels?: boolean;
  onlyWithoutPhoto?: boolean;
  onlyWithPhoto?: boolean;
  parentIds?: string[];
  q?: string;
  fields?: string[];
};

export class FieldsAPI extends BaseAPI {
  getAll() {
    return this.http.get<string[]>({ url: route("/items/fields") });
  }

  getAllValues(field: string) {
    return this.http.get<string[]>({ url: route(`/items/fields/values`, { field }) });
  }
}

export class ItemMaintenanceAPI extends BaseAPI {
  getLog(itemId: string, filters: MaintenanceFilters = {}) {
    return this.http.get<MaintenanceEntryWithDetails[]>({
      url: route(`/items/${itemId}/maintenance`, { status: filters.status?.toString() }),
    });
  }

  create(itemId: string, data: MaintenanceEntryCreate) {
    return this.http.post<MaintenanceEntryCreate, MaintenanceEntry>({
      url: route(`/items/${itemId}/maintenance`),
      body: data,
    });
  }
}

export class ItemsApi extends BaseAPI {
  attachments: GenericAttachmentsAPI;
  maintenance: ItemMaintenanceAPI;
  fields: FieldsAPI;

  constructor(http: Requests, token: string) {
    super(http, token);
    this.fields = new FieldsAPI(http);
    this.attachments = new GenericAttachmentsAPI(http);
    this.maintenance = new ItemMaintenanceAPI(http);
  }

  getAttachments(id: string) {
    return this.attachments.getAll("item", id);
  }

  addAttachment(id: string, file: File, filename: string, type?: AttachmentTypes, primary?: boolean) {
    return this.attachments.add("item", id, file, filename, type, primary);
  }

  updateAttachment(
    attachmentId: string,
    payload: { title?: string; type?: AttachmentTypes; primary?: boolean }
  ) {
    return this.attachments.update(attachmentId, payload);
  }

  deleteAttachment(attachmentId: string) {
    return this.attachments.delete(attachmentId);
  }

  fullpath(id: string) {
    return this.http.get<ItemPath[]>({ url: route(`/items/${id}/path`) });
  }

  getAll(q: ItemsQuery = {}) {
    return this.http.get<ItemSummaryPaginationResult<ItemSummary>>({ url: route("/items", q) });
  }

  create(item: ItemCreate) {
    return this.http.post<ItemCreate, ItemOut>({ url: route("/items"), body: item });
  }

  async get(id: string) {
    const payload = await this.http.get<ItemOut>({ url: route(`/items/${id}`) });

    if (!payload.data) {
      return payload;
    }

    payload.data = parseDate(payload.data, ["purchaseTime", "soldTime", "warrantyExpires", "archivedAt"]);
    return payload;
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/items/${id}`) });
  }

  async update(id: string, item: ItemUpdate) {
    const payload = await this.http.put<ItemCreate, ItemOut>({
      url: route(`/items/${id}`),
      body: this.dropFields(item),
    });
    if (!payload.data) {
      return payload;
    }

    payload.data = parseDate(payload.data, ["purchaseTime", "soldTime", "warrantyExpires", "archivedAt"]);
    return payload;
  }

  async patch(id: string, item: ItemPatch) {
    const resp = await this.http.patch<ItemPatch, ItemOut>({
      url: route(`/items/${id}`),
      body: this.dropFields(item),
    });

    if (!resp.data) {
      return resp;
    }

    resp.data = parseDate(resp.data, ["purchaseTime", "soldTime", "warrantyExpires", "archivedAt"]);
    return resp;
  }

  import(file: File | Blob) {
    const formData = new FormData();
    formData.append("csv", file);

    return this.http.post<FormData, void>({
      url: route("/items/import"),
      data: formData,
    });
  }

  exportURL() {
    return route("/items/export");
  }
}
