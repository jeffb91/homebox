import { BaseAPI, route } from "../base";
import type { ItemOut } from "../types/data-contracts";
import type { AttachmentTypes } from "../types/non-generated";
import type { Requests } from "~~/lib/requests";

export class AttachmentsAPI extends BaseAPI {
  // Haal attachments op voor een generiek relatedType/relatedId
  getAttachments(relatedType: string, relatedId: string) {
    return this.http.get<ItemOut>({
      url: route(`/attachments`, { relatedType, relatedId }),
    });
  }

  addAttachment(
    relatedType: string,
    relatedId: string,
    file: File | Blob,
    filename: string,
    type: AttachmentTypes | null = null,
    primary?: boolean
  ) {
    const formData = new FormData();
    formData.append("file", file);
    if (type) formData.append("type", type);
    formData.append("name", filename);
    if (primary !== undefined) formData.append("primary", primary.toString());

    return this.http.post<FormData, ItemOut>({
      url: route(`/attachments`, { relatedType, relatedId }),
      data: formData,
    });
  }

  deleteAttachment(attachmentId: string) {
    return this.http.delete<void>({
      url: route(`/attachments/${attachmentId}`),
    });
  }

  updateAttachment(attachmentId: string, data: Partial<{ title?: string; primary?: boolean; type?: AttachmentTypes }>) {
    return this.http.put<typeof data, ItemOut>({
      url: route(`/attachments/${attachmentId}`),
      body: data,
    });
  }
}
