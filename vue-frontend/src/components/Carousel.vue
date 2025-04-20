<script setup>
const props = defineProps({
    preview: Array,
});
</script>
<template>
    <!-- image carousel with controls -->
    <div class="flex gap-x-3 overflow-x-auto">
        <div
            v-for="card in preview"
            :key="card.id"
            class="bg-dark h-80 rounded-xl bg-cover text-center"
            :style="`min-width: 241px; min-height: 336px; background-image: url(${card.images.small})`"
        ></div>
    </div>
</template>
