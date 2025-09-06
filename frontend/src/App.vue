<template>
  <div id="app">
    <TodoList>
      <template v-slot:state-slot>
        <StateButton :stateButtons="stateButtons" :activeNumber="activeNumber" @change-show-state="changeShowState" />
      </template>
      <template v-slot:content-slot>
        <ListContent :items="formattedItems" @change-item-state="changeItemState">
          <template v-slot:delete-button-slot="{ itemId }">
            <DeleteItemButton @delete-item="deleteItem" :item-id="itemId"></DeleteItemButton>
          </template>
        </ListContent>
      </template>
      <template v-slot:add-slot>
        <AddSection @add-item="addItem"></AddSection>
      </template>
    </TodoList>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue';
import { todoAPI } from './api';
import TodoList from './components/TodoList.vue';
import StateButton from './components/StateButton.vue';
import DeleteItemButton from './components/DeleteItemButton.vue';
import ListContent from './components/ListContent.vue'
import AddSection from './components/AddItemButton.vue'

const items = ref([])

const stateButtons = [{
  content: "未完成",
}, {
  content: "已完成",
}, {
  content: "全部",
},]

const activeNumber = ref(2)
const formattedItems = computed(() => {
  switch (activeNumber.value) {
    case 0: return items.value.filter(item => !item.isCompleted)  // 未完成
    case 1: return items.value.filter(item => item.isCompleted)   // 已完成
    default: return items.value  // 全部
  }
})

async function fetchItems() {
  try{
    const response = await todoAPI.getItems()
    items.value = response.data||response||[]
    console.log("获取待办事项成功:", items.value)
  } catch (error) {
    console.error("获取待办事项失败:", error)
  }
}

function changeShowState(idx) {
  console.log("show state is changed to", idx)
  activeNumber.value = idx
}

async function changeItemState(id) {
  const item = items.value.find(item => item.id === id)
  if (item) {
    try {
      console.log(item)
      item.isCompleted = !item.isCompleted
      console.log(item)
      await todoAPI.updateItem(item)

      console.log("change item completed state, id:", id, "is completed:", item.isCompleted)
    } catch (error) {
      console.error('failed to update item:', item)
      item.isCompleted = !item.isCompleted
      alert("更新失败，请重试")
    }
  } else {
    console.warn(`Item with id ${id} not found`)
  }
}

async function addItem(content) {
  const newItem = {
    content: content,
    isCompleted: false,
  }

  try {
const createdItem = await todoAPI.addItem(newItem)
    items.value.push(createdItem)
  } catch (error) {
    console.error('failed to add item,content:', content)
    alert("添加失败，请重试")
  }


}

async function deleteItem(id) {
  try {
    //等待调用API删除后端数据
    await todoAPI.deleteItem(id)

    //调用成功后再更新本地数据
    const index = items.value.findIndex(item => item.id === id)
    if (index !== -1) {
      items.value.splice(index, 1)

      console.log("Item deleted, id:", id)
    }
  } catch (error) {
    console.error('failed to delete item,id:', id)
    alert("删除失败，请重试")
  }
}

onMounted(()=>{
  fetchItems()
})

</script>

<style>
:root,
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background-color: #b83b5e;
}
</style>