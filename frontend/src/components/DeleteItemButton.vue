<template>
    <div class="delete-button-container">
        <button class="delete-button" @click="deleteItem">
            X
        </button>
    </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';

const emit = defineEmits(['delete-item'])
const props = defineProps({
    itemId: {
        type: String,
        required: true,
    }
})

function deleteItem() {
    emit('delete-item', props.itemId)
}
</script>

<style scoped>
.delete-button-container {
    display: flex;
    transition: all 1s ease-in-out;
    background-color: #f9aeae;
    width: 100%;
    aspect-ratio: 1;
    border-radius: 100%;
    border: 1px solid red;
    overflow: hidden;
}

/* todo：此处无法正常使用hover和active等 */

.delete-button {
    border: none;
    width: 100%;
    height: 100%;
    font-size: 5vw;
    color: rgb(255, 0, 0);
    background-color: transparent;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 1s ease-in-out;
}
</style>