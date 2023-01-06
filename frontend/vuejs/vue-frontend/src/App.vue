<template>
  <div id="app">
    <ToDos v-bind:todos="todos"/>
  </div>
</template>


<script>
import ToDos from "@/components/ToDos.vue";

export default {
  name: 'app',
  components: {ToDos},

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
  }

}
</script>


<!-- if not scoped, will be applied globally -->
<style scoped>
</style>
