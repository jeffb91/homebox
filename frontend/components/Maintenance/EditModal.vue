<template>
  <Dialog dialog-id="edit-maintenance">
    <DialogContent class="overflow-y-auto max-h-[95vh]">
      <DialogHeader>
        <DialogTitle>
          {{ entry.id ? t("maintenance.modal.edit_title") : t("maintenance.modal.new_title") }}
        </DialogTitle>
      </DialogHeader>

      <form class="flex flex-col gap-2" @submit.prevent="dispatchFormSubmit">
        <!-- Onderhoud velden -->
        <FormTextField v-model="entry.name" autofocus :label="t('maintenance.modal.entry_name')" />
        <DatePicker v-model="entry.completedDate" :label="t('maintenance.modal.completed_date')" />
        <DatePicker v-model="entry.scheduledDate" :label="t('maintenance.modal.scheduled_date')" />
        <FormTextArea v-model="entry.description" :label="t('maintenance.modal.notes')" />
        <FormTextField v-model="entry.measurement" :label="t('maintenance.modal.measurement')" />
        <FormTextField v-model="entry.cost" :label="t('maintenance.modal.cost')" />

        <!-- Attachments sectie -->
        <div class="mt-4">
          <Card ref="attDropZone" class="overflow-visible shadow-xl">
              <div class="px-4 py-5 sm:px-6">
                <h3 class="text-lg font-medium leading-6">{{ $t("items.attachments") }}</h3>
                <p class="text-xs">{{ $t("items.changes_persisted_immediately") }}</p>
              </div>
              <div class="border-t p-4">
                <div v-if="attDropZoneActive" class="grid grid-cols-4 gap-4">
                  <DropZone @drop="dropPhoto"> {{ $t("items.photos") }} </DropZone>
                  <DropZone @drop="dropWarranty"> {{ $t("items.warranty") }} </DropZone>
                  <DropZone @drop="dropManual"> {{ $t("items.manuals") }} </DropZone>
                  <DropZone @drop="dropAttachment"> {{ $t("items.attachments") }} </DropZone>
                  <DropZone @drop="dropReceipt"> {{ $t("items.receipts") }} </DropZone>
                </div>
                <button
                  v-else
                  class="grid h-24 w-full place-content-center border-2 border-dashed border-primary"
                  @click="clickUpload"
                >
                  <input ref="refAttachmentInput" hidden type="file" @change="uploadImage" />
                  <p>{{ $t("items.drag_and_drop") }}</p>
                </button>
              </div>

              <div class="border-t p-4">
                <ul role="list" class="divide-y rounded-md border">
                  <li
                    v-for="attachment in entry.attachments"
                    :key="attachment.id"
                    class="grid grid-cols-6 justify-between py-3 pl-3 pr-4 text-sm"
                  >
                    <p class="col-span-4 my-auto">
                      {{ attachment.filename }}
                    </p>
                    <p class="my-auto">
                      {{ $t(`entry.${attachment.type}`) || "" }}
                    </p>
                    <div class="flex justify-end gap-2">
                      <Tooltip>
                        <TooltipTrigger as-child>
                          <Button variant="destructive" size="icon" @click="deleteAttachment(attachment.id)">
                            <MdiDelete />
                          </Button>
                        </TooltipTrigger>
                        <TooltipContent>{{ $t("global.delete") }}</TooltipContent>
                      </Tooltip>
                      <Tooltip>
                        <TooltipTrigger as-child>
                          <Button size="icon" @click="openAttachmentEditDialog(attachment)">
                            <MdiPencil />
                          </Button>
                        </TooltipTrigger>
                        <TooltipContent>{{ $t("global.edit") }}</TooltipContent>
                      </Tooltip>
                      <Tooltip>
                        <TooltipTrigger as-child>
                          <a
                            :class="buttonVariants({ size: 'icon' })"
                            :href="attachmentURL(attachment.id)"
                            :download="attachment.title"
                          >
                            <MdiDownload />
                          </a>
                        </TooltipTrigger>
                        <TooltipContent> {{ $t("components.item.attachments_list.download") }} </TooltipContent>
                      </Tooltip>
                    </div>
                  </li>
                </ul>
              </div>
            </Card>
        </div>

        <DialogFooter class="mt-4 flex justify-end">
          <Button type="submit">
            <MdiPost />
            {{ entry.id ? t("maintenance.modal.edit_action") : t("maintenance.modal.new_action") }}
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>

  <!-- Bijlage bewerken dialoog -->
  <Dialog dialog-id="attachment-edit">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>{{ t("items.edit.edit_attachment_dialog.title") }}</DialogTitle>
      </DialogHeader>

      <FormTextField
        v-model="editState.title"
        :label="t('items.edit.edit_attachment_dialog.attachment_title')"
      />

      <Label>{{ t("items.edit.edit_attachment_dialog.attachment_type") }}</Label>
      <Select v-model:model-value="editState.type">
        <SelectTrigger>
          <SelectValue :placeholder="t('items.edit.edit_attachment_dialog.select_type')" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem
            v-for="opt in attachmentOpts"
            :key="opt.value"
            :value="opt.value"
          >
            {{ opt.text }}
          </SelectItem>
        </SelectContent>
      </Select>

      <div v-if="editState.type === AttachmentTypes.Photo" class="mt-3 flex items-center gap-2">
        <Checkbox
          id="primary"
          v-model="editState.primary"
          :label="t('items.edit.edit_attachment_dialog.primary_photo')"
        />
      </div>

      <DialogFooter>
        <Button :loading="editState.loading" @click="updateAttachment">
          {{ t("global.update") }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>

  <ul>
    <li v-for="attachment in entry.attachments" :key="attachment.id">
      {{ attachment.filename }}
    </li>
  </ul>
</template>

<script setup lang="ts">
  import { useI18n } from "vue-i18n";
  import { toast } from "@/components/ui/sonner";
  import type {MaintenanceEntry,MaintenanceEntryAttachment, MaintenanceEntryWithDetails } from "~/lib/api/types/data-contracts-old";
  import MdiPost from "~icons/mdi/post";
  import DatePicker from "~~/components/Form/DatePicker.vue";
  import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
  import { useDialog } from "@/components/ui/dialog-provider";

  //imports voor bijlage
  import { reactive, ref } from "vue";
  import { Select, SelectTrigger, SelectContent, SelectItem } from "@/components/ui/select";
  import { Checkbox } from "@/components/ui/checkbox";
  import MdiDelete from "~icons/mdi/delete";
  import MdiPencil from "~icons/mdi/pencil";
  import { AttachmentTypes } from "~~/lib/api/types/non-generated";
  import MdiPaperclip from "~icons/mdi/paperclip";
  import MdiDownload from "~icons/mdi/download";
  import MdiOpenInNew from "~icons/mdi/open-in-new";
  import { buttonVariants } from "@/components/ui/button";
  //import { useDropZone } from "@/composables/useDropZone";
  //import DropZone from "@/components/DropZone.vue";
  //import { Card, Button, Label } from "@/components/ui";
  //import { FormTextField, FormTextArea } from "@/components/Form";
  //import { useConfirm } from "@/composables/useConfirm";

  const { openDialog, closeDialog } = useDialog();

  const { t } = useI18n();
  const api = useUserApi();

  const emit = defineEmits(["changed"]);

  const entry = reactive({
    id: null as string | null,
    name: "",
    completedDate: null as Date | null,
    scheduledDate: null as Date | null,
    description: "",
    cost: "",
    measurement: "",
    itemId: null as string | null,
    attachments: [] as MaintenanceEntryAttachment[],
  });

  async function dispatchFormSubmit() {
    if (entry.id) {
      await editEntry();
      return;
    }

    await createEntry();
  }

   async function createEntry() {
    if (!entry.itemId) {
      return;
    }
    const { error } = await api.items.maintenance.create(entry.itemId, {
      name: entry.name,
      completedDate: entry.completedDate ?? "",
      scheduledDate: entry.scheduledDate ?? "",
      description: entry.description,
      measurement: entry.measurement, 
      cost: parseFloat(entry.cost) ? entry.cost : "0",
    });

    if (error) {
      toast.error(t("maintenance.toast.failed_to_create"));
      return;
    }

    closeDialog("edit-maintenance");
    emit("changed");
  }

  async function editEntry() {
    if (!entry.id) {
      return;
    }

    const { error } = await api.maintenance.update(entry.id, {
      name: entry.name,
      completedDate: entry.completedDate ?? "null",
      scheduledDate: entry.scheduledDate ?? "null",
      description: entry.description,
      measurement: entry.measurement,
      cost: entry.cost,
    });

    if (error) {
      toast.error(t("maintenance.toast.failed_to_update"));
      return;
    }

    closeDialog("edit-maintenance");
    emit("changed");
  }

  const openCreateModal = (itemId: string) => {
    entry.id = null;
    entry.name = "";
    entry.completedDate = null;
    entry.scheduledDate = null;
    entry.description = "";
    entry.cost = "";
    entry.measurement = "";
    entry.measurement = "";
    entry.itemId = itemId;
    openDialog("edit-maintenance");
  };

  const openUpdateModal = async (maintenanceEntry: MaintenanceEntry | MaintenanceEntryWithDetails) => {
    entry.id = maintenanceEntry.id;
    entry.name = maintenanceEntry.name;
    entry.completedDate = new Date(maintenanceEntry.completedDate);
    entry.scheduledDate = new Date(maintenanceEntry.scheduledDate);
    entry.description = maintenanceEntry.description;
    entry.cost = maintenanceEntry.cost;
    entry.measurement = maintenanceEntry.measurement;
    entry.itemId = null;

    const res = await api.maintenance.getAttachments(entry.id ?? "");
    console.log("API response van getAttachments:", res);
    //probleem hieronder is een spookprobleem in studio code -> negeren!
    entry.attachments = res.data.attachments ?? [];
    console.log("entry.attachments na ophalen:", entry.attachments);

    openDialog("edit-maintenance");
  };

  const confirm = useConfirm();

  async function deleteEntry(id: string) {
    const result = await confirm.open(t("maintenance.modal.delete_confirmation"));
    if (result.isCanceled) {
      return;
    }

    const { error } = await api.maintenance.delete(id);

    if (error) {
      toast.error(t("maintenance.toast.failed_to_delete"));
      return;
    }
    emit("changed");
  }

  async function complete(maintenanceEntry: MaintenanceEntry) {
    const { error } = await api.maintenance.update(maintenanceEntry.id, {
      name: maintenanceEntry.name,
      completedDate: new Date(Date.now()),
      scheduledDate: maintenanceEntry.scheduledDate ?? "null",
      description: maintenanceEntry.description,
      measurement: maintenanceEntry.measurement,
      cost: maintenanceEntry.cost,
    });
    if (error) {
      toast.error(t("maintenance.toast.failed_to_update"));
    }
    emit("changed");
  }

  function duplicate(maintenanceEntry: MaintenanceEntry | MaintenanceEntryWithDetails, itemId: string) {
    entry.id = null;
    entry.name = maintenanceEntry.name;
    entry.completedDate = null;
    entry.scheduledDate = null;
    entry.description = maintenanceEntry.description;
    entry.cost = maintenanceEntry.cost;
    entry.measurement = "";
    entry.itemId = itemId;
    openDialog("edit-maintenance");
  }

  defineExpose({ openCreateModal, openUpdateModal, deleteEntry, complete, duplicate });

  //bijlage gerelateerde code
  const attDropZone = ref<HTMLDivElement>();
  const { isOverDropZone: attDropZoneActive } = useDropZone(attDropZone);

  const refAttachmentInput = ref<HTMLInputElement>();

  const MEA = ref<MaintenanceEntryAttachment & { labelIds: string[] }>(null as any);

  function clickUpload() {
    if (!refAttachmentInput.value) {
      return;
    }
    refAttachmentInput.value.click();
  }

  function uploadImage(e: Event) {
    const files = (e.target as HTMLInputElement).files;
    if (!files || !files.item(0)) {
      return;
    }

    const first = files.item(0);
    if (!first) {
      return;
    }

    uploadAttachment([first], null);
  }

  const dropPhoto = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Photo);
  const dropAttachment = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Attachment);
  const dropWarranty = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Warranty);
  const dropManual = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Manual);
  const dropReceipt = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Receipt);

  async function uploadAttachment(files: File[] | null, type: AttachmentTypes | null) {
    if (!files || files.length === 0) {
      return;
    }

    const { data, error } = await api.attachments.add(entry.id ?? "", files[0], files[0].name, type);

    if (error) {
      toast.error(t("items.toast.failed_upload_attachment"));
      return;
    }


    toast.success(t("items.toast.attachment_uploaded"));

    entry.attachments = data.attachments ??[];
    console.log(entry.attachments);
    console.log("API response:", data);
    console.log("entry.attachments:", entry.attachments);
  }

  async function deleteAttachment(attachmentId: string) {
    const confirmed = await confirm.open(t("items.delete_attachment_confirm"));

    if (confirmed.isCanceled) {
      return;
    }

    const { error, data } = await api.attachment.deleteAttachment( attachmentId);
    if (error) {
      toast.error(t("items.toast.failed_delete_attachment"));
      return;
    }
    toast.success(t("items.toast.attachment_deleted"));
    entry.attachments = data.attachments ?? [];
  }

    const editState = reactive({
    loading: false,

    // Values
    obj: {},
    id: "",
    title: "",
    type: "",
    primary: false,
  });

    const attachmentOpts = Object.entries(AttachmentTypes).map(([key, value]) => ({
    text: key[0].toUpperCase() + key.slice(1),
    value,
  }));

    function openAttachmentEditDialog(attachment: MaintenanceEntryAttachment) {
    editState.id = attachment.id;
    editState.title = attachment.filename ?? "";
    editState.type = attachment.type ?? "";
    editState.primary = attachment.primary ?? false;
    openDialog("attachment-edit");

    //editState.obj = attachmentOpts.find(o => o.value === attachment.type) || attachmentOpts[0];
    console.log("Attachment type bij openen:", attachment.type);
  }

    async function updateAttachment() {
    editState.loading = true;
    const { error, data } = await api.maintenance.updateAttachment(editState.id, {
      title: editState.title,
      type: editState.type,
      primary: editState.primary,
    });

    if (error) {
      toast.error(t("items.toast.failed_update_attachment"));
      editState.loading = false;
      return;
    }
    entry.attachments = data.attachments ?? [];

    editState.loading = false;
    editState.id = "";
    editState.title = "";
    editState.type = "";

    toast.success(t("items.toast.attachment_updated"));
  }

  function attachmentURL(attachmentId: string) {
    //return api.authURL(`/maintenance/${props.itemId}/attachments/${attachmentId}`);
    return api.authURL(`/maintenance/attachments/${attachmentId}`); 
  }
  console.log("entry.attachments in template:", entry.attachments);
</script>
