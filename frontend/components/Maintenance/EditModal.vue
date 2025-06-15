<template>
  <Dialog dialog-id="edit-maintenance">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>
          {{ entry.id ? $t("maintenance.modal.edit_title") : $t("maintenance.modal.new_title") }}
        </DialogTitle>
      </DialogHeader>

      <form class="flex flex-col gap-4" @submit.prevent="dispatchFormSubmit">
        <FormTextField v-model="entry.name" :label="$t('maintenance.modal.entry_name')" autofocus />
        <div class="grid grid-cols-2 gap-4">
          <DatePicker v-model="entry.completedDate" :label="$t('maintenance.modal.completed_date')" />
          <DatePicker v-model="entry.scheduledDate" :label="$t('maintenance.modal.scheduled_date')" />
        </div>
        <FormTextArea v-model="entry.description" :label="$t('maintenance.modal.notes')" />
        <FormTextField v-model="entry.measurement" :label="$t('maintenance.modal.measurement')" />
        <FormTextField v-model="entry.cost" :label="$t('maintenance.modal.cost')" />

        <DialogFooter class="mt-2">
          <Button type="submit">
            <MdiPost />
            {{ entry.id ? $t("maintenance.modal.edit_action") : $t("maintenance.modal.new_action") }}
          </Button>
        </DialogFooter>
      </form>

      <!-- Attachments Section -->
      <section class="mt-6 border-t pt-4">
        <h3 class="text-lg font-medium">{{ $t("maintenance.attachments.title") }}</h3>

        <div ref="dropzone" class="mt-2 border-2 border-dashed rounded-lg p-6 text-center"
             :class="{ 'border-primary bg-primary/10': isOverDropZone }"
             @dragover.prevent @drop.prevent="onDrop">
          {{ $t("maintenance.attachments.drag_and_drop") }}
        </div>

        <div class="mt-4 flex gap-2">
          <input id="file-input" hidden type="file" multiple @change="onFileSelect" />
          <label for="file-input">
            <Button size="sm">{{ $t("maintenance.attachments.select_files") }}</Button>
          </label>
          <Button size="sm" :disabled="!files.length || uploading" @click="uploadFiles">
            {{ $t("maintenance.attachments.upload_button") }}
          </Button>
        </div>

        <ul v-if="attachments.length" class="mt-4 divide-y border rounded-md">
          <li v-for="att in attachments" :key="att.id" class="flex items-center justify-between p-3">
            <span class="truncate">{{ att.title || att.filename }}</span>
            <div class="flex gap-2">
              <Button size="icon" @click="downloadAttachment(att)">⬇️</Button>
              <Button size="icon" variant="destructive" @click="deleteAttachment(att.id)">🗑️</Button>
              <Button size="icon" @click="openEditAttachment(att)">✏️</Button>
            </div>
          </li>
        </ul>
      </section>

      <!-- Edit Attachment Dialog -->
      <Dialog dialog-id="edit-attachment">
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{{ $t("maintenance.attachments.edit_title") }}</DialogTitle>
          </DialogHeader>
          <div class="flex flex-col gap-4">
            <FormTextField v-model="editState.title" :label="$t('maintenance.attachments.field_title')" />
          </div>
          <DialogFooter>
            <Button @click="updateAttachment" :disabled="editState.loading">{{ $t("global.update") }}</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from "vue";
import { useDialog } from "@/components/ui/dialog-provider";
import { useDropZone } from "@/components/hooks/useDropZone";
import { FormTextField } from "@/components/Form/TextField.vue";
import DatePicker from "@/components/Form/DatePicker.vue";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import MdiPost from "~icons/mdi/post";
import { useI18n } from "vue-i18n";
import { toast } from "@/components/ui/sonner";
import type { MaintenanceEntry, MaintenanceEntryAttachment } from "~~/lib/api/types/data-contracts";
const api = useUserApi();
const { t } = useI18n();
const { openDialog, closeDialog } = useDialog();

const entry = reactive<Partial<MaintenanceEntry>>({});
const files = ref<File[]>([]);
const attachments = ref<MaintenanceEntryAttachment[]>([]);
const uploading = ref(false);

const editState = reactive({ id: "", title: "", loading: false });

const dropzone = ref<HTMLElement>();
const { isOverDropZone } = useDropZone(dropzone);

function dispatchFormSubmit() {
  entry.id ? editEntry() : createEntry();
}

async function createEntry() {
  if (!entry.itemId) return;
  const { error } = await api.items.maintenance.create(entry.itemId, {
    name: entry.name,
    completedDate: entry.completedDate ?? "",
    scheduledDate: entry.scheduledDate ?? "",
    description: entry.description,
    measurement: entry.measurement,
    cost: parseFloat(entry.cost) ? entry.cost : "0",
  });

  if (error) return toast.error(t("maintenance.toast.failed_to_create"));

  closeDialog("edit-maintenance");
  emit("changed");
}

async function editEntry() {
  if (!entry.id) return;
  const { error } = await api.maintenance.update(entry.id, {
    name: entry.name,
    completedDate: entry.completedDate ?? "",
    scheduledDate: entry.scheduledDate ?? "",
    description: entry.description,
    measurement: entry.measurement,
    cost: entry.cost,
  });

  if (error) return toast.error(t("maintenance.toast.failed_to_update"));

  closeDialog("edit-maintenance");
  emit("changed");
}

watch(() => entry.id, async id => {
  if (id) {
    const res = await api.maintenance.getAttachments(id);
    if (!res.error) attachments.value = res.data;
  }
});

function onDrop(ev: DragEvent) {
  const dt = ev.dataTransfer;
  if (dt?.files) files.value = Array.from(dt.files);
}

function onFileSelect(ev: Event) {
  const input = ev.target as HTMLInputElement;
  if (input.files) files.value = Array.from(input.files);
}

async function uploadFiles() {
  if (!entry.id || !files.value.length) return;
  uploading.value = true;
  for (const file of files.value) {
    const form = new FormData();
    form.append("file", file);
    form.append("name", file.name);
    const { error, data } = await api.maintenance.uploadAttachments(entry.id!, form);
    if (error) {
      toast.error(t("maintenance.attachments.upload_failed"));
      break;
    }
    attachments.value.push(data);
  }
  files.value = [];
  uploading.value = false;
}

function downloadAttachment(att: MaintenanceEntryAttachment) {
  api.maintenance.getAttachmentUrl(att.id)
    .then(url => {
      const a = document.createElement("a");
      a.href = url;
      a.download = att.filename;
      document.body.appendChild(a);
      a.click();
      a.remove();
    })
    .catch(() => toast.error(t("maintenance.attachments.download_failed")));
}

async function deleteAttachment(id: string) {
  if (!confirm(t("maintenance.attachments.delete_confirm"))) return;
  const { error } = await api.maintenance.deleteAttachment(id);
  if (!error) {
    attachments.value = attachments.value.filter(a => a.id !== id);
    toast.success(t("maintenance.attachments.delete_success"));
  } else toast.error(t("maintenance.attachments.delete_failed"));
}

function openEditAttachment(att: MaintenanceEntryAttachment) {
  editState.id = att.id;
  editState.title = att.title || "";
  openDialog("edit-attachment");
}

async function updateAttachment() {
  editState.loading = true;
  const { error } = await api.maintenance.updateAttachment(editState.id, { title: editState.title });
  editState.loading = false;
  if (!error) {
    const att = attachments.value.find(a => a.id === editState.id);
    if (att) att.title = editState.title;
    closeDialog("edit-attachment");
    toast.success(t("maintenance.attachments.updated"));
  } else toast.error(t("maintenance.attachments.update_failed"));
}
</script>
