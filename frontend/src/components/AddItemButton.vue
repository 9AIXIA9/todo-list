<template>
    <div class="add-container">
        <div></div>
        <input type="text" v-model="inputContent" placeholder="请输入事项内容" class="input-section">
        <button @click="addItem" type="submit" class="submit-button">
            新建
        </button>
    </div>
</template>

<script setup>
import { defineEmits, ref } from 'vue';


const emit = defineEmits(['add-item'])

const inputContent = ref("")

function addItem() {
    if (inputContent.value != "") {
        console.log("add the item:" + inputContent.value)
        emit('add-item', inputContent.value)
        inputContent.value = ""
    }
}
</script>

<style scoped>
.add-container {
    display: flex;
    justify-content: space-around;
    padding: 1%;
}

.input-section {
    border-radius: 30%;
    background-color: transparent;
    border: 1vw dotted #b83b5e;
    outline: 1vw solid #ff7c7c;
    padding: 2%;
    color: rgb(0, 0, 0);
    font-size: 3vw;
}

.submit-button {
    transition: all 0.1s ease-in-out;
    border-radius: 30%;
    padding: 1%;
    border: 1px solid #34fff8;
    background-color: #b83b5e;
    color: rgb(0, 0, 0);
    font-size: 3vw;
}

.submit-button:hover{
    transition: all 0.5s ease-in-out;
    background-color: #e3295e;
    padding: 4%;
    cursor: pointer;
}

.submit-button:active{
    transition: all 0.2s ease-in-out;
    padding: 0;
    font-size: 0;
    background-color: #ff729a;
    border: none;
}
</style>