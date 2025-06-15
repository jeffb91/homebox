<template>
  <Dialog dialog-id="edit-maintenance">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>
          {{ entry.id ? $t("maintenance.modal.edit_title") : $t("maintenance.modal.new_title") }}
        </DialogTitle>
      </DialogHeader>

      <form class="flex flex-col gap-2" @submit.prevent="dispatchFormSubmit">
        <FormTextField v-model="entry.name" autofocus :label="$t('maintenance.modal.entry_name')" />
        <DatePicker v-model="entry.completedDate" :label="$t('maintenance.modal.completed_date')" />
        <DatePicker v-model="entry.scheduledDate" :label="$t('maintenance.modal.scheduled_date')" />
        <FormTextArea v-model="entry.description" :label="$t('maintenance.modal.notes')" />
        <FormTextField v-model="entry.measurement" :label="$t('maintenance.modal.measurement')" />
        <FormTextField v-model="entry.cost" autofocus :label="$t('maintenance.modal.cost')" />



        <DialogFooter>
          <Button type="submit">
            <MdiPost />
            {{ entry.id ? $t("maintenance.modal.edit_action") : $t("maintenance.modal.new_action") }}
          </Button>
        </DialogFooter>
      </form>
      <section class="mt-4">
        <h3>{{ $t("maintenance.attachments.title") }}</h3>
        <input type="file" multiple @change="onFileChange" />
        <button @click="uploadFiles" :disabled="uploading || files.length === 0">
          {{ $t("maintenance.attachments.upload_button") }}
        </button>
        <ul>
          <li v-for="attachment in attachments" :key="attachment.id" class="flex items-center justify-between space-x-2">
            <span>{{ attachment.title || attachment.filename }}</span>
            <div class="flex space-x-2">
              <!-- Download knop -->
              <button
                type="button"
                @click="downloadAttachment(attachment)"
                class="btn btn-sm btn-outline"
                :title="$t('maintenance.attachments.download')"
              >
                ⬇️
              </button>

              <!-- Verwijder knop -->
              <button
                type="button"
                @click="deleteAttachment(attachment.id)"
                class="btn btn-sm btn-danger"
                :title="$t('maintenance.attachments.delete')"
              >
                🗑️
              </button>
            </div>
          </li>
        </ul>

      </section>

    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
  import { useI18n } from "vue-i18n";
  import { toast } from "@/components/ui/sonner";
  import type { MaintenanceEntry, MaintenanceEntryWithDetails } from "~~/lib/api/types/data-contracts";
  import MdiPost from "~icons/mdi/post";
  import DatePicker from "~~/components/Form/DatePicker.vue";
  import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
  import { useDialog } from "@/components/ui/dialog-provider";
  import { ref, reactive, watch, onMounted } from "vue";

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
    measurement: "", // ✅ nieuw veld
    itemId: null as string | null,
  });

  const files = ref<File[]>([]);
  const attachments = ref<Array<{ id: string; title?: string; filename: string }>>([]);
  const uploading = ref(false);

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
    entry.itemId = itemId;
    openDialog("edit-maintenance");
  };

  const openUpdateModal = (maintenanceEntry: MaintenanceEntry | MaintenanceEntryWithDetails) => {
    entry.id = maintenanceEntry.id;
    entry.name = maintenanceEntry.name;
    entry.completedDate = new Date(maintenanceEntry.completedDate);
    entry.scheduledDate = new Date(maintenanceEntry.scheduledDate);
    entry.description = maintenanceEntry.description;
    entry.cost = maintenanceEntry.cost;
    entry.measurement = maintenanceEntry.measurement;
    entry.itemId = null;
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
    entry.measurement = null;
    entry.itemId = itemId;
    openDialog("edit-maintenance");
  }

  defineExpose({ openCreateModal, openUpdateModal, deleteEntry, complete, duplicate });

  // Ophalen bestaande attachments als je een entry hebt met id
  async function fetchAttachments() {
    if (!entry.id) return;
    const { data, error } = await api.maintenance.getAttachments(entry.id);
    if (!error) {
      attachments.value = data;
    }
  }

  watch(() => entry.id, () => {
    if (entry.id) fetchAttachments();
  });

  function onFileChange(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files) {
      files.value = Array.from(input.files);
    }
  }

  async function uploadFiles() {
    if (!entry.id || files.value.length === 0) return;
    uploading.value = true;

    const formData = new FormData();
    files.value.forEach(f => formData.append("files", f));

    const { error, data } = await api.maintenance.uploadAttachments(entry.id, formData);
    uploading.value = false;

    if (error) {
      toast.error(t("maintenance.attachments.upload_failed"));
      return;
    }

    // Voeg nieuwe attachments toe aan lijst en reset file input
    attachments.value.push(...data);
    files.value = [];
  }

  async function downloadAttachment(attachment: { id: string; filename: string }) {
    // Simpelste manier: url van backend ophalen en een link forceren om te downloaden
    try {
      // Vervang met jouw backend url die het bestand serveert, bv:
      const url = await api.maintenance.getAttachmentUrl(attachment.id);
      const link = document.createElement("a");
      link.href = url;
      link.download = attachment.filename;
      link.target = "_blank";
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    } catch {
      toast.error(t("maintenance.attachments.download_failed"));
    }
  }

  async function deleteAttachment(id: string) {
    if (!confirm(t("maintenance.attachments.delete_confirm"))) {
      return;
    }
    const { error } = await api.maintenance.deleteAttachment(id);
    if (error) {
      toast.error(t("maintenance.attachments.delete_failed"));
      return;
    }
    // Verwijder attachment lokaal uit lijst
    attachments.value = attachments.value.filter(a => a.id !== id);
    toast.success(t("maintenance.attachments.delete_success"));
  }

</script>

