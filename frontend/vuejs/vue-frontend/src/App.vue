<template>
  <div id="app">
    <ToDos v-bind:todos="todos" v-on:delete-todo="deleteToDos"/>
    <!--  v-on defines events   -->
    <AddToDo v-on:add-todo="addToDo" />
  </div>
</template>


<script>
import ToDos from "@/components/ToDos.vue";
import AddToDo from "./components/AddToDo";

export default {
  name: 'app',
  components: {AddToDo, ToDos},

  data() {
    return {
      todos: [
        {
          id: 1,
          title: 'Go workout',
          completed: false
        }
      ]
    }
  },

  methods: {
    addToDo(newToDo) {
      this.todos = [...this.todos, newToDo];
    },

    deleteToDos(id) {
      this.todos = this.todos.filter(todo => todo.id !== id);
    }
  }

}
</script>


<!-- if not scoped, will be applied globally -->
<style scoped>
</style>
