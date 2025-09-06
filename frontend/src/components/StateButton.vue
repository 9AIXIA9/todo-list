<template>
    <div class="state-button-container">
        <button v-for="(button, index) in props.stateButtons" :key="button.content"
            :class="{ active: index == activeNumber }" @click="changeShowState(index)" class="state-button">
            {{ button.content }}
        </button>
    </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';


const emit = defineEmits(['change-show-state'])

const props = defineProps({
    stateButtons: {
        type: Array,
        required: true
    },
    activeNumber: {
        type: Number,
        required: true,
    }
})

function changeShowState(index) {
    console.log("change state to", index)
    emit('change-show-state', index)
}
</script>

<style scoped>
.state-button-container {
    display: flex;
    justify-content: space-around;
}

.state-button {
    transition: all 1s ease-in-out;
    background-color: #FFF93D;
    border-radius: 50%;
    width: 20%;
    color: rgb(0, 0, 0);
    font-size: 3vw;
    border: 1px solid;
}

.state-button-container :hover {
    background-color: #ffbe26;
    transition: all 0.2s ease-in-out;
    cursor: pointer;
    color: rgb(167, 167, 167);
}

.active {
    background-color: #FF7700;
}
</style>