<template>
    <div class="clock">
        {{ today }}
    </div>
</template>

<script setup>
let now = new Date();
let year = now.getFullYear();
let month = now.getMonth() + 1; // getMonth() 返回0-11，需要加1
let day = now.getDate();

// 格式化月和日，确保两位数表示
month = month < 10 ? '0' + month : month;
day = day < 10 ? '0' + day : day;

const today = year + '-' + month + '-' + day;
</script>

<style scoped>
.clock{
    transition: all 1s ease-in-out;
}
</style>