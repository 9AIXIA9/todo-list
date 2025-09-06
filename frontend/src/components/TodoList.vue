<template>
    <div class="list-container">
        <div class="header">
            <div class="title">Todo List</div>
            <DateClock></DateClock>
        </div>
        <div class="state-section">
            <slot name="state-slot">
                <div>state</div>
            </slot>
        </div>
        <div class="content-section">
            <slot name="content-slot">
                <div>content</div>
            </slot>
        </div>
        <div class="footer">
            <div class="add-section">
                <slot name="add-slot">
                    add item
                </slot>
            </div>
        </div>
    </div>
</template>

<script setup>
import DateClock from './DateClock.vue';

</script>

<style scoped>
* {
    color: rgb(255, 247, 0);
    text-shadow: 1px 1px 3px #ff0000;
    font-size: 3vw;
    background-color: #00FF77;
}

.title {
    transition: all 1s ease-in-out;
}

.list-container {
    margin: 20px auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 80%;
    box-sizing: border-box;
    border: 2px solid black;
    border-radius: 30% 30% 0 0;
    overflow: hidden;
    text-align: center;
    box-shadow: 0 0 0 2px #34fff8;
}

.header {
    transition: all 1s ease-in-out;
    border-bottom: 2px solid rgb(0, 0, 0);
    border-radius: 50px 50px 0 0;
}

.state-section {
    border-bottom: 2px solid rgb(0, 0, 0);
}


.footer {
    transition: all 1s ease-in-out;
    border-top: 2px solid rgb(0, 0, 0);
}
</style>