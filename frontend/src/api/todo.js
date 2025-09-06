import request from "./request";

export const todoAPI = {
  //获取所有待办事项
  getItems() {
    return request.get("/items");
  },

  //添加待办事项
  addItem(item) {
    return request.post("/item", item);
  },

  //更新待办事项
  updateItem(item) {
    return request.put("/item", item);
  },

  //删除待办事项
  deleteItem(id) {
    return request.delete("/item/" + id, id);
  },
};
