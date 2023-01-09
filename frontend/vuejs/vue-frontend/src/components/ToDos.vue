<template>
  <div>
    <h2>My todo list</h2>

    <ul>
      <li v-bind:key="todo.id" v-for="todo in todos">
        <ToDo v-bind:todo="todo" v-on:delete-todo="$emit('delete-todo', todo.id)"/>
      </li>
    </ul>
  </div>
</template>

<script>
import ToDo from "@/components/ToDo.vue";

export default {
  name: "to-dos",
  components: {
    ToDo,
  },

  props: [
    "todos"
  ]
}
</script>

<style scoped>

</style>
