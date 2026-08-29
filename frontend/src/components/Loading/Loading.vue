<script lang="ts" setup>
import ProgressSpinner from 'primevue/progressspinner'
</script>

<template>
    <ProgressSpinner 
        class="stone-spinner"
        style="width: 30px; height: 30px" 
        strokeWidth="8" 
        fill="transparent"
        animationDuration=".8s"
        aria-label="Loading" 
    />
</template>

<style scoped>
.stone-spinner :deep(.p-progress-spinner-circle) {
    stroke: #292524;
    animation: p-progress-spinner-color 6s ease-in-out infinite;
}

@keyframes p-progress-spinner-color {
    100%, 0% {
        stroke: #292524;
    }
    40% {
        stroke: #292524;
    }
    66% {
        stroke: #292524;
    }
    80%, 90% {
        stroke: #292524;
    }
}
</style>