<script setup lang="ts">
  import type { ItemOut } from "~/lib/api/types/data-contracts-old";

  const props = defineProps<{
    item: ItemOut;
  }>();
</script>

<template>
  <BaseContainer class="flex flex-col gap-8">
    <MaintenanceListView :current-item-id="props.item.id"></MaintenanceListView>
  </BaseContainer>
</template>
