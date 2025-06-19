import { BaseAPI } from "../base";
import type { Attachment } from "../types/data-contracts";
import type { AttachmentTypes } from "../types/non-generated";

export class AttachmentsAPI extends BaseAPI {
  // Haal attachments op voor een generiek relatedType/relatedId
getAttachments(relatedType: string, relatedId: string) {
  return this.http.get<Attachment[]>({
    url: `/attachments?related_type=${relatedType}&related_id=${relatedId}`,
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
    formData.append("name", filename);
    formData.append("related_type", relatedType);
    formData.append("related_id", relatedId);
    if (type) formData.append("type", type);
    if (primary !== undefined) formData.append("primary", primary.toString());

    return this.http.post<FormData, Attachment>({
      url: `/attachments`,
      data: formData,
    });
  }

  deleteAttachment(attachmentId: string) {
    return this.http.delete<void>({
      url: `/attachments/${attachmentId}`,
    });
  }

  updateAttachment(attachmentId: string, data: Partial<{ title?: string; primary?: boolean; type?: AttachmentTypes }>) {
    return this.http.put<typeof data, Attachment>({
      url: `/attachments/${attachmentId}`,
      body: data,
    });
  }
}
