<template>
  <div v-bind:class="{ 'completed': todo.completed }">
    <p v-on:click="markComplete">{{ todo.title }}</p>
    <button @click="$emit('delete-todo', todo.id)">Delete</button>
  </div>

</template>

<script>

export default {
  name: "to-do",
  props: [
      "todo"
  ],

  methods: {
    markComplete() {
      // mutating props is kind of bad
      // eslint-disable-next-line vue/no-mutating-props
      this.todo.completed = !this.todo.completed
    }
  }
}
</script>

<style scoped>
  .completed {
    text-decoration: line-through;
  }
</style>
