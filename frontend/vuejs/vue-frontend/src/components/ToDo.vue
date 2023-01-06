<template>
  <div>
    <p>{{ todo.title }}</p>
  </div>

</template>

<script>

export default {
  name: "to-do",
  props: [
      "todo"
  ]
}
</script>

<style scoped>

</style>
