<template>
    <div class="content-container">
        <div v-for="item in props.items" :key="item.id" class="item-container">
            <input @change="changeItemState(item.id)" :checked="item.isCompleted" type="checkbox"
                :id="`checkbox-${item.id}`" name="checkBox" class="invisible">
            <div>

            </div>
            <label :for="`checkbox-${item.id}`">
                <div class="item" :class="{ 'completed-text': item.isCompleted }">
                    {{ item.content }}
                </div>
            </label>
            <div class="delete-button">
                <slot name="delete-button-slot" :itemId="item.id">delete</slot>
            </div>
        </div>
    </div>
</template>

<script setup>
import { defineProps, defineEmits } from "vue"

const props = defineProps({
    items: {
        type: Array,
        required: true
    },
})

const emit = defineEmits(['change-item-state'])

function changeItemState(id) {
    emit('change-item-state', id)
}
</script>

<style scoped>
.content-container {
    display: flex;
    flex-direction: column;
    transition: all 1s ease-in-out;
}

.item-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 70%;
    border-bottom: 2px solid black;
}

.item-container:last-child {
    border-bottom: none;
}

.completed-text {
    text-decoration: line-through wavy 20% rgba(255, 97, 97, 0.534);
}

.invisible{
    display: none;
}

</style>