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
            class="min-h-[350px] min-w-[250px] rounded-xl bg-contain text-center"
            :style="`background-image: url(${card.images.small})`"
        ></div>
    </div>
</template>
